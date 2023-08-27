package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetPollsReq struct {
	g.Meta       `path:"/poll" tags:"Poll" method:"get" summary:"获取投票列表"`
	StatusStart  int                         `json:"status_start" dc:"起始状态" v:"required|min:0"`
	StatusEnd    int                         `json:"status_end" dc:"结束状态" v:"required|max:300"`
	WithRecords  bool                        `json:"with_records" dc:"是否返回投票记录" v:"boolean" d:"false"`
	PolledFilter consts.PollListPolledFilter `json:"polled_filter" dc:"投票过滤器" v:"in:0,1,2" d:"0"`
	Order        string                      `json:"order" dc:"排序方式" v:"in:desc,asc" d:"asc"`
	Page         int                         `json:"page" dc:"页码" v:"min:1" d:"1"`
	PageSize     int                         `json:"page_size" dc:"每页数量" v:"min:0|max:1000" d:"30"`
}

type GetPollsRes model.GetPollListOutput
