package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	bizctx := service.BizCtx().Get(ctx)
	if bizctx == nil || bizctx.User == nil { // 正常情况下不会出现
		g.Log().Error(ctx, service.BizCtx().Get(ctx))
		err = gerror.NewCode(gcode.CodeUnknown, "bizctx or bizctx.User is nil")
	}
	user := bizctx.User
	res = &v1.GetUserRes{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Poll: v1.UserPoll{
			Points: v1.UserPollPoints{
				Total:      user.Poll.Points,
				Approve:    user.Poll.Accept,
				Reject:     user.Poll.Reject,
				NeedModify: user.Poll.NeedEdited,
			},
			Count:     user.Poll.Points / gconv.Int(service.User().GetUserPollPointsByUserRole(ctx, user.Role)), // TODO: 稍后换成数据库 Count
			Score:     user.Poll.Score,
			CreatedAt: user.Poll.CreatedAt,
			UpdatedAt: user.Poll.UpdatedAt,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return
}
