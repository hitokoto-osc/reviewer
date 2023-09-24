package hitokoto

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/api/hitokoto/adminV1"
)

func (c *ControllerAdminV1) DeleteList(ctx context.Context, req *adminV1.DeleteListReq) (res *adminV1.DeleteListRes, err error) {
	err = service.Hitokoto().Delete(ctx, req.UUIDs)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	return nil, nil
}
