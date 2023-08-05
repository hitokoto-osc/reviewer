package poll

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) NewPoll(ctx context.Context, req *v1.NewPollReq) (res *v1.NewPollRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
