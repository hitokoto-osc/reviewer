package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserPollResult(ctx context.Context, req *v1.GetUserPollResultReq) (res *v1.GetUserPollResultRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}