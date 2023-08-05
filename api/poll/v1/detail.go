package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type GetPollDetailReq struct {
	g.Meta       `path:"/poll/{id}" tags:"Poll" method:"get" summary:"获取投票详情"`
	SentenceUUID string `json:"sentence_uuid" dc:"句子 UUID" v:"required|length:36"`
}

type PollRecord struct {
	UserID    int64             `json:"user_id" dc:"用户 ID"`
	Point     int               `json:"point" dc:"投票点数"`
	Type      consts.PollMethod `json:"type" dc:"投票类型"`
	Comment   string            `json:"comment" dc:"理由"`
	CreatedAt string            `json:"created_at" dc:"投票时间"`
	UpdatedAt string            `json:"updated_at" dc:"更新时间"`
}

type GetPollDetailRes struct {
	PollSchema
	Records []PollRecord `json:"logs" dc:"投票记录"`
}
