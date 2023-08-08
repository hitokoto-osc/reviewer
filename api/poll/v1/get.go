package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetPollsReq struct {
	g.Meta       `path:"/poll" tags:"Poll" method:"get" summary:"获取投票列表"`
	StatusStart  int  `json:"status_start" dc:"起始状态" v:"required|min:0"`
	StatusEnd    int  `json:"status_end" dc:"结束状态" v:"required|max:300"`
	WithPollData bool `json:"polled" dc:"是否返回投票信息" v:"boolean" d:"false"`
}

type GetPollsRes []struct {
	model.PollElement
	PollData model.PollData `json:"poll_data" dc:"投票数据"`
}
