package poll

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) GetPolls(ctx context.Context, req *v1.GetPollsReq) (res *v1.GetPollsRes, err error) {
	user := service.BizCtx().GetUser(ctx) // User must not be nil
	if req.StatusStart > req.StatusEnd {
		req.StatusStart, req.StatusEnd = req.StatusEnd, req.StatusStart
	}
	if user.Role != consts.UserRoleAdmin && req.StatusEnd >= int(consts.PollStatusApproved) {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "权限不足")
	}
	keys, e := g.DB().GetCache().KeyStrings(ctx)
	g.Log().Debugf(ctx, "%+v %+v", keys, e)
	out, err := service.Poll().GetPollList(ctx, &model.GetPollListInput{
		StatusStart:     req.StatusStart,
		StatusEnd:       req.StatusEnd,
		Order:           req.Order,
		UserID:          user.Id,
		WithPollRecords: req.WithRecords,
		WithMarks:       true,
		WithCache:       true,
		PolledFilter:    req.PolledFilter,
		Page:            req.Page,
		PageSize:        req.PageSize,
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "获取投票列表失败")
	}
	res = (*v1.GetPollsRes)(out)
	// keys, e = g.DB().GetCache().KeyStrings(ctx)
	// g.Log().Debugf(ctx, "%+v %+v", keys, e)
	return res, nil
}
