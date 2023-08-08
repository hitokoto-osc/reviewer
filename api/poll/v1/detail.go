package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetPollDetailReq struct {
	g.Meta       `path:"/poll/detail" tags:"Poll" method:"get" summary:"获取投票详情"`
	SentenceUUID string `json:"sentence_uuid" dc:"句子 UUID" v:"required|size:36#句子 UUID 必须填写|UUID 长度有误" in:"query"`
}

type GetPollDetailRes struct {
	model.PollElement
	Records []model.PollRecord `json:"logs" dc:"投票记录"`
}
