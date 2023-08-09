package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type PollReq struct {
	g.Meta  `path:"/poll/:id" tags:"Poll" method:"put" summary:"投票"`
	ID      uint              `json:"id" dc:"投票ID" v:"required|uint" in:"path"`
	Method  consts.PollMethod `json:"method" dc:"投票方式" v:"required|enums"`
	MarkID  uint              `json:"mark_id" dc:"标记" v:"required-unless:method,1"`
	Comment string            `json:"comment" dc:"理由" v:"required-if:method,2,3|length:1,255"`
}

// PollRes 成功返回句子的投票记录
type PollRes = GetPollDetailRes
