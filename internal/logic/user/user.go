package user

import (
	"context"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/hitokoto-osc/reviewer/internal/dao"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// VerifyAPIV1Token 用于 v1 接口校验用户是否登录
// TODO: v2 中会使用新的用户系统，并且将会使用带有 ACL、签名的授权机制。目前的 token 机制会被废弃。
func (s *sUser) VerifyAPIV1Token(ctx context.Context, token string) (bool, error) {
	user, err := dao.Users.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour, // 缓存一小时
		Name:     "user:token:" + token,
		Force:    false,
	}).Where("token = ?", token).One()
	if err != nil || user == nil {
		return false, err
	}
	return true, nil
}
