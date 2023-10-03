package notification

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

const (
	SentenceReviewedNotificationExchange   = "notification"
	SentenceReviewedNotificationRoutingKey = "notification.hitokoto_reviewed"
)

var SentenceReviewedNotificationSettingField = dao.UserNotification.Columns().EmailNotificationHitokotoReviewed

type SentenceReviewedNotificationMessage struct {
	To         string  `json:"to"`          // Email
	UUID       string  `json:"uuid"`        // SentenceUUID
	Hitokoto   string  `json:"hitokoto"`    // Sentence
	From       string  `json:"from"`        // From
	FromWho    *string `json:"from_who"`    // FromWho
	Type       string  `json:"type"`        // Hitokoto Type
	Creator    string  `json:"creator"`     // Creator
	CreatorUID uint    `json:"creator_uid"` // CreatorUID
	CreatedAt  string  `json:"created_at"`  // Sentence CreatedAt

	OperatedAt   string `json:"operated_at"`   // 操作时间
	ReviewerName string `json:"reviewer_name"` // 审核员名称
	ReviewerUID  int    `json:"reviewer_uid"`  // 审核员用户标识
	Status       int    `json:"status"`        // 审核结果： 200 为通过，201 为驳回
}

// SentenceReviewedNotification 发送句子审核通知
func (s *sNotification) SentenceReviewedNotification(
	ctx context.Context,
	poll *model.PollElement,
	reviewerUID uint,
	reviewerName string) error {
	ok, err := s.IsUserShouldDoNotification(ctx, SentenceReviewedNotificationSettingField, poll.Sentence.CreatorUID)
	if err != nil {
		return err
	} else if !ok {
		return nil // 用户没有开启通知
	}
	user, err := service.User().GetUserByID(ctx, poll.Sentence.CreatorUID)
	if err != nil {
		return gerror.Wrap(err, "获取用户信息失败")
	}
	data := make([]any, 0, 1)
	data = append(data, SentenceReviewedNotificationMessage{
		To:           user.Email,
		UUID:         poll.SentenceUUID,
		Hitokoto:     poll.Sentence.Hitokoto,
		From:         poll.Sentence.From,
		FromWho:      poll.Sentence.FromWho,
		Type:         string(poll.Sentence.Type),
		Creator:      poll.Sentence.Creator,
		CreatorUID:   poll.Sentence.CreatorUID,
		CreatedAt:    poll.Sentence.CreatedAt,
		OperatedAt:   (*gtime.Time)(poll.UpdatedAt).Format("c"),
		ReviewerName: reviewerName,
		ReviewerUID:  int(reviewerUID),
		Status:       int(poll.Status),
	})
	return DoNotification(ctx, SentenceReviewedNotificationExchange, SentenceReviewedNotificationRoutingKey, data)
}
