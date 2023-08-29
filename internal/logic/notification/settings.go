package notification

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/dao"
)

func (s *sNotification) GetUserIDsShouldDoNotification(
	ctx context.Context,
	userIDs []uint,
	settingField string,
) ([]uint, error) {
	records, err := dao.UserNotification.Ctx(ctx).
		Fields(dao.UserNotification.Columns().UserId).                    // 只获取 UserId
		Where(dao.UserNotification.Columns().UserId, userIDs).            // 给定 UserIds 列表
		Where(dao.UserNotification.Columns().EmailNotificationGlobal, 1). // 全局开启邮件通知
		Where(settingField, 1).                                           // 该权限开启通知
		Array()
	if err != nil {
		return nil, err
	}
	return gconv.Uints(records), nil
}

func (s *sNotification) GetUsersShouldDoNotification(
	ctx context.Context,
	settingField string,
) ([]entity.Users, error) {
	users, err := service.User().GetReviewersAndAdmins(ctx)
	if err != nil {
		return nil, err
	}
	usersIDs := make([]uint, 0, len(users))
	for i := 0; i < len(users); i++ {
		usersIDs = append(usersIDs, users[i].Id)
	}
	userIdsThatShouldDoNotification, err := s.GetUserIDsShouldDoNotification(
		ctx,
		usersIDs,
		PollCreatedNotificationSettingField,
	)
	if err != nil {
		return nil, err
	}
	if len(userIdsThatShouldDoNotification) == 0 {
		return nil, nil // 没有用户需要发送通知
	}
	// filter users
	tmp := make([]entity.Users, 0, len(userIdsThatShouldDoNotification))
	for _, userID := range userIdsThatShouldDoNotification {
		for i := 0; i < len(users); i++ {
			if users[i].Id == userID {
				tmp = append(tmp, users[i])
			}
		}
	}
	users = tmp
	return users, nil
}

func (s *sNotification) IsUserShouldDoNotification(
	ctx context.Context,
	settingField string,
	userID uint,
) (bool, error) {
	record, err := dao.UserNotification.Ctx(ctx).
		Where(dao.UserNotification.Columns().UserId, userID).             // 给定 UserIds 列表
		Where(dao.UserNotification.Columns().EmailNotificationGlobal, 1). // 全局开启邮件通知
		Where(settingField, 1).                                           // 该权限开启通知
		One()
	if err != nil {
		return false, err
	}
	return record != nil, nil
}
