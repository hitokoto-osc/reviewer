package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserPollUnreviewed(ctx context.Context, req *v1.GetUserPollUnreviewedReq) (
	res *v1.GetUserPollUnreviewedRes,
	err error,
) {
	user := service.BizCtx().GetUser(ctx)
	count, err := service.Poll().CountUserUnreviewedPoll(ctx, user.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取未审查的投票数量失败")
	}
	return &v1.GetUserPollUnreviewedRes{
		Count: count,
	}, nil
}
