package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type PollReq struct {
	g.Meta       `path:"/poll" tags:"Poll" method:"put" summary:"投票"`
	Method       consts.PollMethod `json:"method" dc:"投票方式" v:"required|enums"`
	SentenceUUID string            `json:"sentence_uuid" dc:"句子 UUID" v:"required|length:36"`
	MarkID       uint              `json:"mark_id" dc:"标记" v:"required-unless:method,1"`
	Comment      string            `json:"comment" dc:"理由" v:"required-if:method,2,3|length:1,255"`
}

// PollRes 成功返回句子的投票记录
type PollRes struct {
	model.PollElement
	PollData model.PollData `json:"poll_data" dc:"投票数据"`
}
