package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type ModifySentenceReq struct {
	g.Meta       `path:"/admin/sentence/{uuid}" tags:"Admin" method:"put" summary:"修改句子"`
	SentenceUUID string                     `json:"sentence_uuid" dc:"句子 UUID" v:"required|length:36"`
	Params       model.ModifySentenceParams `json:"params" dc:"句子参数"`
}

type ModifySentenceRes struct {
	model.PollSchema
}
