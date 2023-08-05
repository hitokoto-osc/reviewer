// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type (
	IUser interface {
		GetRoleByUser(ctx context.Context, user *entity.Users) (role consts.UserRole, err error)
		MustGetRoleByUser(ctx context.Context, user *entity.Users) (role consts.UserRole)
		GetUserStatusByUser(ctx context.Context, user *entity.Users) (status consts.UserStatus, err error)
		MustGetUserStatusByUser(ctx context.Context, user *entity.Users) (status consts.UserStatus)
		GetUserPollPointsByUserRole(ctx context.Context, role consts.UserRole) (points consts.UserPollPoints)
		// VerifyAPIV1Token 用于 v1 接口校验用户是否登录
		// TODO: v2 中会使用新的用户系统，并且将会使用带有 ACL、签名的授权机制。目前的 token 机制会被废弃。
		VerifyAPIV1Token(ctx context.Context, token string) (flag bool, err error)
		GetUserByToken(ctx context.Context, token string) (user *entity.Users, err error)
		GetUserByID(ctx context.Context, id uint) (user *entity.Users, err error)
		GetPollUserByUserID(ctx context.Context, id uint) (user *entity.PollUsers, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
