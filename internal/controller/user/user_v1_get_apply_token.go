package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetApplyToken(ctx context.Context, req *v1.GetApplyTokenReq) (res *v1.GetApplyTokenRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
