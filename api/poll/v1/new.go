package v1

import "github.com/gogf/gf/v2/frame/g"

type NewPollReq struct {
	g.Meta `path:"/poll" tags:"Poll" method:"post" summary:"新建投票"`
}

type NewPollRes struct {
	PollSchema
	RemainPending int `json:"remain_pending" dc:"剩余待处理"`
}
