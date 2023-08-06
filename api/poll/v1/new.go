package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type NewPollReq struct {
	g.Meta `path:"/poll" tags:"Poll" method:"post" summary:"新建投票"`
}

type NewPollRes struct {
	model.PollSchema
	RemainPending int `json:"remain_pending" dc:"剩余待处理"`
}
