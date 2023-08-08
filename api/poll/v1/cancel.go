package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type CancelPollReq struct {
	g.Meta       `path:"/poll/{id}" tags:"Poll" method:"delete" summary:"取消投票"`
	SentenceUUID string `json:"sentence_uuid" dc:"句子 UUID" v:"required|length:36"`
}

type CancelPollRes struct {
	model.PollElement
	PollData model.PollData `json:"poll_data" dc:"投票数据"`
}
