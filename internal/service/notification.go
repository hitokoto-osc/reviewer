// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type (
	INotification interface {
		PollCreatedNotification(ctx context.Context, poll *model.PollElement) error
		// DailyReportNotification 发送每日报告通知
		// 直接交给 Event 生产消息
		DailyReportNotification(ctx context.Context, rawData []model.DailyReportNotificationMessage) error
		PollFinishedNotification(ctx context.Context, poll *model.PollElement, pollLogs []entity.PollLog) error
		// SentenceReviewedNotification 发送句子审核通知
		SentenceReviewedNotification(ctx context.Context, poll *model.PollElement, reviewerUID uint, reviewerName string) error
		GetUserIDsShouldDoNotification(ctx context.Context, userIDs []uint, settingField string) ([]uint, error)
		GetUsersShouldDoNotification(ctx context.Context, settingField string) ([]entity.Users, error)
		IsUserShouldDoNotification(ctx context.Context, settingField string, userID uint) (bool, error)
	}
)

var (
	localNotification INotification
)

func Notification() INotification {
	if localNotification == nil {
		panic("implement not found for interface INotification, forgot register?")
	}
	return localNotification
}

func RegisterNotification(i INotification) {
	localNotification = i
}
