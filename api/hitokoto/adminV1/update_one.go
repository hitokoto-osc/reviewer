package adminV1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type UpdateOneReq struct {
	g.Meta `path:"/hitokoto/:uuid" method:"PUT" summary:"更新句子" tags:"句子"`
	UUID   string `json:"uuid" dc:"句子 UUID" v:"required|regex:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$#请输入 UUID|UUID 非法" in:"path"` //nolint:lll
	// 句子相关
	model.DoHitokotoV1Update
}

type UpdateOneRes struct{}
