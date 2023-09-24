package hitokoto

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/hitokoto/v1"
)

func (c *ControllerV1) GetHitokoto(ctx context.Context, req *v1.GetHitokotoReq) (res *v1.GetHitokotoRes, err error) {
	hitokoto, err := service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, req.UUID)
	if err != nil {
		return nil, gerror.Wrap(err, "获取一言失败")
	}
	if hitokoto == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "一言不存在")
	}
	res = (*v1.GetHitokotoRes)(hitokoto)
	return
}
