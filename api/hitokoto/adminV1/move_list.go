package adminV1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type MoveListReq struct {
	g.Meta `path:"/hitokoto/move" method:"POST" summary:"移动句子" tags:"句子"`
	UUIDs  []string              `json:"uuids" dc:"句子 UUID" v:"required|length:1,100"`
	Target consts.HitokotoStatus `json:"target" dc:"目标状态" v:"required|in:pending,approved,rejected" in:"query"`
}

type MoveListRes struct {
	IsSuccess   bool        `json:"is_success"`
	Total       int         `json:"total"`
	FailedUUIDs []string    `json:"failed_uuids"`
	FailedDesc  g.MapStrStr `json:"failed_desc"`
}
