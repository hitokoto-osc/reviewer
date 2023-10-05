package hitokoto

import (
	"context"
	"strings"

	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"github.com/samber/lo"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/api/hitokoto/adminV1"
)

func (c *ControllerAdminV1) GetList(ctx context.Context, req *adminV1.GetListReq) (res *adminV1.GetListRes, err error) {
	// fmt.Printf("GetList: %+v", req)
	if len(req.UUIDs) > 0 {
		req.UUIDs = lo.Map(req.UUIDs, func(item string, index int) string {
			return strings.Trim(item, "\" '") // 去除首尾的引号
		})
	}
	out, err := service.Hitokoto().GetList(ctx, &model.GetHitokotoV1SchemaListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		UUIDs:    req.UUIDs,
		Keywords: req.Keywords,
		Creator:  req.Creator,
		From:     req.From,
		FromWho:  req.FromWho,
		Type:     req.Type,
		Status:   req.Status,
		Start:    req.Start,
		End:      req.End,
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取句子列表失败")
	}
	res = (*adminV1.GetListRes)(out)
	return
}
