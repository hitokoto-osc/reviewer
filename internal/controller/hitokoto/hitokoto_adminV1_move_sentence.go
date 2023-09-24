package hitokoto

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/api/hitokoto/adminV1"
)

func (c *ControllerAdminV1) MoveSentence(ctx context.Context, req *adminV1.MoveSentenceReq) (res *adminV1.MoveSentenceRes, err error) {
	sentence, err := service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, req.UUID)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	if sentence == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "未找到指定的句子")
	}
	err = service.Hitokoto().Move(ctx, sentence, req.Target)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
