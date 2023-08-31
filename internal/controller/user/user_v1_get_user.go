package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error) {
	// 从 BizCtx 中获取用户信息
	user := service.BizCtx().GetUser(ctx)
	if user == nil { // 正常情况下不会出现
		g.Log().Error(ctx, service.BizCtx().Get(ctx))
		err = gerror.NewCode(gcode.CodeUnknown, "bizctx.User is nil")
		return nil, err
	}

	var pollLogs []model.UserPollLogWithSentenceAndUserMarks
	// 是否需要附带投票记录
	if req.WithPollLogs {
		var out *model.GetUserPollLogsWithSentenceOutput
		out, err = service.User().GetUserPollLogsWithSentence(ctx, model.GetUserPollLogsInput{
			UserID:    user.Id,
			Order:     dao.PollLog.Columns().CreatedAt + " desc",
			Page:      1,
			PageSize:  30,
			WithCache: true,
		})
		if err != nil {
			return nil, err
		}
		pollLogs = out.Collection
	}
	count := user.Poll.Points / int(service.User().GetUserPollPointsByUserRole(user.Role)) // TODO: 稍后如果更换点数的话，务必换成数据库 Count

	res = &v1.GetUserRes{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Poll: model.UserPoll{
			Points: model.UserPollPoints{
				Total:      user.Poll.Points,
				Approve:    user.Poll.Accept,
				Reject:     user.Poll.Reject,
				NeedModify: user.Poll.NeedEdited,
			},
			Count:        count,
			Score:        user.Poll.Score,
			AdoptionRate: user.Poll.AdoptionRate,
			CreatedAt:    (*time.Time)(user.Poll.CreatedAt),
			UpdatedAt:    (*time.Time)(user.Poll.UpdatedAt),
		},
		PollLog:   pollLogs,
		CreatedAt: (*time.Time)(user.CreatedAt),
		UpdatedAt: (*time.Time)(user.UpdatedAt),
	}
	return res, nil
}
