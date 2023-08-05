package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type GetUserPollLogReq struct {
	g.Meta `path:"/user/poll/log" tags:"User" method:"get" summary:"获取用户投票记录"`
	Offset int `json:"offset" dc:"偏移量" v:"required|integer#偏移量必须为整数"`
	Limit  int `json:"limit" dc:"数量" v:"required|integer#数量必须为整数"`
}

type UserPollLog struct {
	Point        int               `json:"point" dc:"投票点数"`
	SentenceUUID string            `json:"sentence_uuid" dc:"句子 UUID"`
	Type         consts.PollStatus `json:"type" dc:"投票类型"`
	Comment      string            `json:"comment" dc:"理由"`
	CreatedAt    string            `json:"created_at" dc:"投票时间"`
	UpdatedAt    string            `json:"updated_at" dc:"更新时间"`
}

type GetUserPollLogRes []UserPollLog
