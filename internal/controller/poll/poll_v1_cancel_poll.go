package poll

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) CancelPoll(ctx context.Context, req *v1.CancelPollReq) (res *v1.CancelPollRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
