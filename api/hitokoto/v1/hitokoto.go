package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetHitokotoReq struct {
	g.Meta `path:"/hitokoto/:uuid" method:"get" summary:"获取一言" tags:"Hitokoto"`
	UUID   string `json:"uuid" dc:"一言 UUID" v:"required|regex:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$#请输入 UUID|UUID 非法" in:"path"` //nolint:lll
}

type GetHitokotoRes model.HitokotoV1WithPoll
