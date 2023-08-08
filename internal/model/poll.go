package model

import (
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/utility/time"
)

type PollRecord struct {
	UserID    uint              `json:"user_id" dc:"用户 ID"`
	Point     int               `json:"point" dc:"投票点数"`
	Type      consts.PollMethod `json:"type" dc:"投票类型"`
	Comment   string            `json:"comment" dc:"理由"`
	CreatedAt *time.Time        `json:"created_at" dc:"投票时间"`
	UpdatedAt *time.Time        `json:"updated_at" dc:"更新时间"`
}

type PollMark struct {
	ID        uint                    `json:"id" dc:"标记 ID"`
	Text      string                  `json:"text" dc:"标记文本"`
	Level     consts.PollMarkLevel    `json:"level" dc:"标记等级"`
	Property  consts.PollMarkProperty `json:"property" dc:"标记属性"`
	UpdatedAt *time.Time              `json:"updated_at" dc:"更新时间"`
	CreatedAt *time.Time              `json:"created_at" dc:"创建时间"`
}

type PollData struct {
	Point  int               `json:"point" dc:"投票点数"`
	Method consts.PollMethod `json:"method" dc:"投票方式"`
}

type PollElement struct {
	SentenceUUID       string            `json:"sentence_uuid" dc:"句子 UUID"`
	Sentence           HitokotoV1Schema  `json:"sentence" dc:"句子"`
	Status             consts.PollStatus `json:"status" dc:"投票状态"`
	Accept             int               `json:"accept" dc:"赞同票数"`
	Reject             int               `json:"reject" dc:"反对票数"`
	NeedModify         int               `json:"need_edited" dc:"需要修改票数"`
	NeedCommonUserPoll int               `json:"need_common_user_poll" dc:"需要普通用户投票"`
	CreatedAt          *time.Time        `json:"created_at" dc:"创建时间"`
	UpdatedAt          *time.Time        `json:"updated_at" dc:"更新时间"`
}
