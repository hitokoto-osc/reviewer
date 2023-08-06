package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserPollLog(ctx context.Context, req *v1.GetUserPollLogReq) (res *v1.GetUserPollLogRes, err error) {
	records, err := service.User().GetUserPollLogsWithSentence(ctx, model.UserPollLogsInput{
		Order:     dao.PollLog.Columns().CreatedAt + " " + req.Order,
		Page:      req.Page,
		PageSize:  req.PageSize,
		WithCache: true,
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取用户投票记录失败")
	}
	res = &v1.GetUserPollLogRes{
		UserPollLogsWithSentenceOutput: *records,
	}
	return res, nil
}
