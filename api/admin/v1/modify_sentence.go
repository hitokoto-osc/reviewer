package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

type SentenceParams struct {
	Hitokoto string `json:"hitokoto" dc:"句子内容" v:"length:1,1000#句子长度应该在 1-1000 之间"`
	Type     string `json:"type" dc:"句子类型" v:"length:1,100#句子类型长度应该在 1-100 之间"`
	From     string `json:"from" dc:"句子来源" v:"length:1,100#来源长度应该在 1-100 之间"`
	FromWho  string `json:"from_who" dc:"句子来源者" v:"length:1,100#来源者长度应该在 1-100 之间"`
	Creator  string `json:"creator" dc:"句子创建者" v:"length:1,100#创建者长度应该在 1-100 之间"`
}

type ModifySentenceReq struct {
	g.Meta       `path:"/admin/sentence/{uuid}" tags:"Admin" method:"put" summary:"修改句子"`
	SentenceUUID string         `json:"sentence_uuid" dc:"句子 UUID" v:"required|length:36"`
	Params       SentenceParams `json:"params" dc:"句子参数"`
}

type ModifySentenceRes struct {
	v1.PollSchema
}
