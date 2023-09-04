package user

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/model/do"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"

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
func (s *sUser) VerifyAPIV1Token(ctx context.Context, token string) (flag bool, err error) {
	var user *entity.Users
	user, err = s.GetUserByToken(ctx, token)
	if err != nil || user == nil {
		return
	}
	if userStatus, _ := s.GetUserStatusByUser(ctx, user); userStatus != consts.UserStatusNormal { // User 必定非空，因此不用判断错误
		return
	}
	flag = true
	return
}

func (s *sUser) GetUserByToken(ctx context.Context, token string) (user *entity.Users, err error) {
	err = dao.Users.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour, // 缓存一小时
		Name:     "user:token:" + token,
		Force:    false,
	}).Where(do.Users{Token: token}).Scan(&user)
	return
}

func (s *sUser) GetUserByID(ctx context.Context, id uint) (user *entity.Users, err error) {
	err = dao.Users.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour, // 缓存一小时
		Name:     "user:id:" + gconv.String(id),
		Force:    false,
	}).Where(do.Users{Id: id}).Scan(&user)
	return
}

func (s *sUser) SetUserRoleReviewer(ctx context.Context, userID uint) error {
	var userToken string
	if userID <= 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return gerror.New("user not found")
		}
		userID = user.Id
		userToken = user.Token
	} else {
		user, err := s.GetUserByID(ctx, userID)
		if err != nil {
			return err
		}
		userToken = user.Token
	}
	res, err := dao.Users.Ctx(ctx).
		Data(dao.Users.Columns().IsReviewer, 1).
		Where(dao.Users.Columns().Id, userID).
		Update()
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return gerror.New("failed to update user")
	}
	// clear cache
	if _, err = g.DB().GetCache().Remove(ctx, "SelectCache:user:token:"+userToken, "SelectCache:user:id:"+gconv.String(userID)); err != nil {
		g.Log().Error(ctx, err)
	}
	return nil
}
