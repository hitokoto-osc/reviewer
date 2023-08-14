package notification

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/hitokoto-osc/reviewer/internal/dao"
)

const (
	PollDailyReportNotificationExchange   = "notification"
	PollDailyReportNotificationRoutingKey = "notification.hitokoto_poll_daily_report"
)

var PollDailyReportNotificationSettingField = dao.UserNotification.Columns().EmailNotificationPollDailyReport

// DailyReportNotification 发送每日报告通知
// 直接交给 Event 生产消息
func (s *sNotification) DailyReportNotification(ctx context.Context, rawData []model.DailyReportNotificationMessage) error {
	users, err := s.GetUsersShouldDoNotification(ctx, PollDailyReportNotificationSettingField)
	if err != nil {
		return err
	}
	// filter
	data := make([]any, 0, len(users))
	for i := 0; i < len(users); i++ {
		v := &users[i]
		for j := 0; j < len(rawData); j++ {
			if rawData[j].To == v.Email {
				data = append(data, rawData[j])
			}
		}
	}
	return DoNotification(ctx, PollDailyReportNotificationExchange, PollDailyReportNotificationRoutingKey, data)
}
