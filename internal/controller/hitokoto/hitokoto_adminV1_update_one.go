package hitokoto

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/mitchellh/mapstructure"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/api/hitokoto/adminV1"
)

func (c *ControllerAdminV1) UpdateOne(ctx context.Context, req *adminV1.UpdateOneReq) (res *adminV1.UpdateOneRes, err error) {
	sentence, err := service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, req.UUID)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	if sentence == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "未找到指定的句子")
	}

	var do = make(g.Map)
	err = mapstructure.Decode(req.DoHitokotoV1Update, &do)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	if _, ok := do["from_who"]; !ok && ghttp.RequestFromCtx(ctx).GetBody() != nil {
		var content = make(g.Map)
		err = json.Unmarshal(ghttp.RequestFromCtx(ctx).GetBody(), &content)
		if err != nil {
			return nil, gerror.WrapCode(gcode.CodeInternalError, err)
		}
		if _, ok := content["from_who"]; ok && req.DoHitokotoV1Update.FromWho == nil {
			do["from_who"] = nil
		}
	}

	fmt.Printf("do: %+v", do)

	err = service.Hitokoto().UpdateByUUID(ctx, sentence, do)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
