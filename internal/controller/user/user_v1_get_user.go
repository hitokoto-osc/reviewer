package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserRes, error) {
	var err error
	// 从 BizCtx 中获取用户信息
	bizctx := service.BizCtx().Get(ctx)
	if bizctx == nil || bizctx.User == nil { // 正常情况下不会出现
		g.Log().Error(ctx, service.BizCtx().Get(ctx))
		err = gerror.NewCode(gcode.CodeUnknown, "bizctx or bizctx.User is nil")
		return nil, err
	}
	user := bizctx.User

	var pollLogs []model.UserPollLog
	var count int
	// 是否需要附带投票记录
	if req.WithPollLogs {
		var pollLogsList []entity.PollLog
		pollLogsList, err = service.User().GetUserPollLogByUserID(ctx, user.Id)
		if err != nil {
			return nil, err
		}
		count = len(pollLogsList)
		pollLogs = make([]model.UserPollLog, count)
		for i, v := range pollLogsList {
			pollLogs[i] = model.UserPollLog{
				Point:        v.Point,
				SentenceUUID: v.SentenceUuid,
				Type:         consts.PollStatus(v.Type),
				Comment:      v.Comment,
				CreatedAt:    (*time.Time)(v.CreatedAt),
				UpdatedAt:    (*time.Time)(v.UpdatedAt),
			}
		}
	} else {
		count = user.Poll.Points / int(service.User().GetUserPollPointsByUserRole(user.Role)) // TODO: 稍后如果更换点数的话，务必换成数据库 Count
	}

	res := &v1.GetUserRes{
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
			Count:     count,
			Score:     user.Poll.Score,
			CreatedAt: (*time.Time)(user.Poll.CreatedAt),
			UpdatedAt: (*time.Time)(user.Poll.UpdatedAt),
		},
		PollLog:   pollLogs,
		CreatedAt: (*time.Time)(user.CreatedAt),
		UpdatedAt: (*time.Time)(user.UpdatedAt),
	}
	return res, nil
}
