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
	data := make([]any, 0, len(rawData))
	for i := 0; i < len(rawData); i++ {
		data = append(data, rawData[i])
	}
	return DoNotification(ctx, PollDailyReportNotificationExchange, PollDailyReportNotificationRoutingKey, data)
}
