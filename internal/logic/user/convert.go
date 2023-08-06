package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

func (s *sUser) GetRoleByUser(ctx context.Context, user *entity.Users) (role consts.UserRole, err error) {
	if user == nil {
		return consts.UserRoleGuest, gerror.New("user is nil")
	}
	role = consts.UserRoleUser
	if user.IsReviewer == 1 {
		role = consts.UserRoleReviewer
	}
	if user.IsAdmin == 1 {
		role = consts.UserRoleAdmin
	}
	return
}

func (s *sUser) MustGetRoleByUser(ctx context.Context, user *entity.Users) (role consts.UserRole) {
	var err error
	role, err = s.GetRoleByUser(ctx, user)
	if err != nil {
		g.Log("service.user.MustGetRoleByUser").Panic(ctx, err)
	}
	return
}

func (s *sUser) GetUserStatusByUser(ctx context.Context, user *entity.Users) (status consts.UserStatus, err error) {
	if user == nil {
		return consts.UserStatusSuspended, gerror.New("user is nil")
	}
	status = consts.UserStatusNormal
	if user.IsSuspended == 1 {
		status = consts.UserStatusSuspended
	}
	return
}

func (s *sUser) MustGetUserStatusByUser(ctx context.Context, user *entity.Users) (status consts.UserStatus) {
	var err error
	status, err = s.GetUserStatusByUser(ctx, user)
	if err != nil {
		g.Log("service.user.MustGetUserStatusByUser").Panic(ctx, err)
	}
	return
}

func (s *sUser) GetUserPollPointsByUserRole(role consts.UserRole) (points consts.UserPollPoints) {
	points = consts.UserPollPointsNormal
	if role == consts.UserRoleReviewer {
		points = consts.UserPollPointsReviewer
	}
	if role == consts.UserRoleAdmin {
		points = consts.UserPollPointsAdmin
	}
	return
}

func (s *sUser) GetUserRoleCodeByUserRole(role consts.UserRole) (code consts.UserRoleCode) {
	code = consts.UserRoleCodeGuest
	if role == consts.UserRoleUser {
		code = consts.UserRoleCodeUser
	}
	if role == consts.UserRoleReviewer {
		code = consts.UserRoleCodeReviewer
	}
	if role == consts.UserRoleAdmin {
		code = consts.UserRoleCodeAdmin
	}
	return
}
