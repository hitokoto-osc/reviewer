package notification

import (
	"context"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

const (
	PollCreatedNotificationExchange   = "notification"
	PollCreatedNotificationRoutingKey = "notification.hitokoto_poll_created"
)

var PollCreatedNotificationSettingField = dao.UserNotification.Columns().EmailNotificationPollCreated

type PollCreatedNotificationMessage struct {
	To         string  `json:"to"`          // Email
	UUID       string  `json:"uuid"`        // SentenceUUID
	Hitokoto   string  `json:"hitokoto"`    // Sentence
	From       string  `json:"from"`        // From
	FromWho    *string `json:"from_who"`    // FromWho
	Type       string  `json:"type"`        // Hitokoto Type
	Creator    string  `json:"creator"`     // Creator
	CreatorUID uint    `json:"creator_uid"` // CreatorUID
	UserName   string  `json:"user_name"`   // receiver name
	ID         int     `json:"id"`          // Poll ID
	CreatedAt  string  `json:"created_at"`  // Poll CreatedAt ISO 时间
}

func (s *sNotification) PollCreatedNotification(ctx context.Context, poll *model.PollElement) error {
	users, err := s.GetUsersShouldDoNotification(ctx, PollCreatedNotificationSettingField)
	if err != nil {
		return err
	}
	data := make([]any, 0, len(users))
	for i := 0; i < len(users); i++ {
		v := &users[i]
		data = append(data, PollCreatedNotificationMessage{
			To:         v.Email,
			UUID:       poll.SentenceUUID,
			Hitokoto:   poll.Sentence.Hitokoto,
			From:       poll.Sentence.From,
			FromWho:    poll.Sentence.FromWho,
			Type:       string(poll.Sentence.Type),
			Creator:    poll.Sentence.Creator,
			CreatorUID: poll.Sentence.CreatorUID,
			UserName:   v.Name,
			ID:         int(poll.ID),
			CreatedAt:  (*gtime.Time)(poll.CreatedAt).Format("c"),
		})
	}
	return DoNotification(ctx, PollCreatedNotificationExchange, PollCreatedNotificationRoutingKey, data)
}
