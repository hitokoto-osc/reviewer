package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetOnePollsReq struct {
	g.Meta         `path:"/hitokoto/{uuid}/polls" method:"get" summary:"获取一言最新的投票信息" tags:"Hitokoto"`
	UUID           string `json:"uuid" dc:"一言 UUID" v:"required|regex:^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$#请输入 UUID|UUID 非法" in:"path"` //nolint:lll
	WithPolledData bool   `json:"with_polled_data" dc:"是否返回投票记录" v:"boolean" d:"false"`
	Page           int    `json:"page" dc:"页码" v:"min:1" d:"1"`
	PageSize       int    `json:"page_size" dc:"每页数量" v:"min:0|max:1000" d:"30"`
}

type GetOnePollsRes model.Page[model.PollListElement]
