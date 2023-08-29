package notification

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

const (
	PollFinishedNotificationExchange   = "notification"
	PollFinishedNotificationRoutingKey = "notification.hitokoto_poll_finished"
)

var PollFinishedNotificationSettingField = dao.UserNotification.Columns().EmailNotificationPollResult

type PollFinishedNotificationMessage struct {
	To         string  `json:"to"`          // Email
	UUID       string  `json:"uuid"`        // SentenceUUID
	Hitokoto   string  `json:"hitokoto"`    // Sentence
	From       string  `json:"from"`        // From
	FromWho    *string `json:"from_who"`    // FromWho
	Type       string  `json:"type"`        // Hitokoto Type
	Creator    string  `json:"creator"`     // Creator
	CreatorUID uint    `json:"creator_uid"` // CreatorUID

	ID        int    `json:"id"`         // 投票 ID
	UpdatedAt string `json:"updated_at"` // 投票更新时间，这里也是结束时间
	UserName  string `json:"user_name"`  // 审核员名字
	CreatedAt string `json:"created_at"` // 投票创建时间
	Status    int    `json:"status"`     // 投票结果： 200 入库，201 驳回，202 需要修改
	Method    int    `json:"method"`     // 审核员投票方式： 1 入库，2 驳回，3 需要修改
	Point     int    `json:"point"`      // 审核员投的票数
	// TODO: 加入审核员投票的意见标签？
}

func (s *sNotification) PollFinishedNotification(
	ctx context.Context,
	poll *model.PollElement,
	pollLogs []entity.PollLog,
) error {
	usersThatShouldDoNotification, err := s.GetUsersShouldDoNotification(ctx, PollFinishedNotificationSettingField)
	if err != nil {
		return err
	}
	// 取交集
	if len(usersThatShouldDoNotification) == 0 || len(pollLogs) == 0 {
		return nil
	}
	data := make([]any, 0, len(pollLogs))
	for i := 0; i < len(usersThatShouldDoNotification); i++ {
		for j := 0; j < len(pollLogs); j++ {
			if usersThatShouldDoNotification[i].Id == uint(pollLogs[j].UserId) {
				data = append(data, PollFinishedNotificationMessage{
					To:         usersThatShouldDoNotification[i].Email,
					UUID:       poll.SentenceUUID,
					Hitokoto:   poll.Sentence.Hitokoto,
					From:       poll.Sentence.From,
					FromWho:    poll.Sentence.FromWho,
					Type:       string(poll.Sentence.Type),
					Creator:    poll.Sentence.Creator,
					CreatorUID: poll.Sentence.CreatorUID,

					ID:        pollLogs[j].PollId,
					UpdatedAt: pollLogs[j].UpdatedAt.Format("c"),
					UserName:  usersThatShouldDoNotification[i].Name,
					CreatedAt: pollLogs[j].CreatedAt.Format("c"),
					Status:    int(poll.Status),
					Method:    pollLogs[j].Type,
					Point:     pollLogs[j].Point,
				})
			}
		}
	}
	return DoNotification(ctx, PollFinishedNotificationExchange, PollFinishedNotificationRoutingKey, data)
}
