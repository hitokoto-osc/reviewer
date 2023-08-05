package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/admin/v1"
)

func (c *ControllerV1) ModifySentence(ctx context.Context, req *v1.ModifySentenceReq) (res *v1.ModifySentenceRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
