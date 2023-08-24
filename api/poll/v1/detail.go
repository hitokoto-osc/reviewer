package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetPollDetailReq struct {
	g.Meta         `path:"/poll/:id" tags:"Poll" method:"get" summary:"获取投票详情"`
	ID             int  `json:"id" dc:"投票ID" v:"required|min:1" in:"path"`
	WithPolledData bool `json:"with_polled_data" dc:"是否返回投票记录" v:"boolean" d:"false"`
}

type GetPollDetailRes struct {
	model.PollListElement
}
