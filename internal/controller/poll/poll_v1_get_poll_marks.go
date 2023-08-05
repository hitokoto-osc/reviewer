package poll

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) GetPollMarks(ctx context.Context, req *v1.GetPollMarksReq) (res *v1.GetPollMarksRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
