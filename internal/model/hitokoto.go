package model

import "github.com/hitokoto-osc/reviewer/internal/consts"

type ModifySentenceParams struct {
	Hitokoto string `json:"hitokoto" dc:"句子内容" v:"length:1,1000#句子长度应该在 1-1000 之间"`
	Type     string `json:"type" dc:"句子类型" v:"length:1,100#句子类型长度应该在 1-100 之间"`
	From     string `json:"from" dc:"句子来源" v:"length:1,100#来源长度应该在 1-100 之间"`
	FromWho  string `json:"from_who" dc:"句子来源者" v:"length:1,100#来源者长度应该在 1-100 之间"`
	Creator  string `json:"creator" dc:"句子创建者" v:"length:1,100#创建者长度应该在 1-100 之间"`
}

type HitokotoV1Schema struct {
	ID         uint                  `json:"id" dc:"句子 ID"`
	UUID       string                `json:"uuid" dc:"句子 UUID"`
	Hitokoto   string                `json:"hitokoto" dc:"句子内容"`
	Type       consts.HitokotoType   `json:"type" dc:"句子类型"`
	From       string                `json:"from" dc:"句子来源"`
	FromWho    *string               `json:"from_who" dc:"句子来源者"`
	Creator    string                `json:"creator" dc:"句子创建者"`
	CreatorUID uint                  `json:"creator_uid" dc:"句子创建者 ID"`
	Reviewer   uint                  `json:"reviewer" dc:"句子审查者 ID"`
	Status     consts.HitokotoStatus `json:"status" dc:"句子状态"`
	PollStatus consts.PollStatus     `json:"poll_status" dc:"句子投票状态"`
	CreatedAt  string                `json:"created_at" dc:"创建时间"`
}
