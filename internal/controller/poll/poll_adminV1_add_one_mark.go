package poll

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/api/poll/adminV1"
)

func (c *ControllerAdminV1) AddOneMark(ctx context.Context, req *adminV1.AddOneMarkReq) (res *adminV1.AddOneMarkRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
