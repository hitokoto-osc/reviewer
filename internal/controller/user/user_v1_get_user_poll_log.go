package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserPollLog(ctx context.Context, req *v1.GetUserPollLogReq) (res *v1.GetUserPollLogRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
