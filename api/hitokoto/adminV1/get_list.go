package adminV1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

// GetListReq 获取句子列表
type GetListReq struct {
	g.Meta `path:"/hitokoto" method:"GET" summary:"获取句子列表" tags:"句子"`
	model.GetHitokotoV1SchemaListInput
}

type GetListRes model.GetHitokotoV1SchemaListOutput
