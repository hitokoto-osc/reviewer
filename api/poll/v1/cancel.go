package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CancelPollReq struct {
	g.Meta `path:"/poll/:id/cancel" tags:"Poll" method:"delete" summary:"取消投票"`
	ID     int `json:"id" dc:"投票 ID" v:"required|min:1#请输入投票 ID"`
}

type CancelPollRes struct{}
