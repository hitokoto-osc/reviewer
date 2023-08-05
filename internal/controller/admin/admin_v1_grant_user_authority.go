package admin

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/admin/v1"
)

func (c *ControllerV1) GrantUserAuthority(ctx context.Context, req *v1.GrantUserAuthorityReq) (res *v1.GrantUserAuthorityRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
