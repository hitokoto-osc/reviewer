package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type ModifySentenceParams struct {
	Hitokoto string `json:"hitokoto" dc:"句子内容" v:"length:1,1000#句子长度应该在 1-1000 之间"`
	Type     string `json:"type" dc:"句子类型" v:"length:1,100#句子类型长度应该在 1-100 之间"`
	From     string `json:"from" dc:"句子来源" v:"length:1,100#来源长度应该在 1-100 之间"`
	FromWho  string `json:"from_who" dc:"句子来源者" v:"length:1,100#来源者长度应该在 1-100 之间"`
	Creator  string `json:"creator" dc:"句子创建者" v:"length:1,100#创建者长度应该在 1-100 之间"`
}

type HitokotoV1 struct {
	ID         uint                  `json:"id" dc:"句子 ID"`
	UUID       string                `json:"uuid" dc:"句子 UUID"`
	Hitokoto   string                `json:"hitokoto" dc:"句子内容"`
	Type       consts.HitokotoType   `json:"type" dc:"句子类型"`
	From       string                `json:"from" dc:"句子来源"`
	FromWho    *string               `json:"from_who" dc:"句子来源者"`
	Creator    string                `json:"creator" dc:"句子创建者"`
	CreatorUID uint                  `json:"creator_uid" dc:"句子创建者 ID"`
	Reviewer   uint                  `json:"reviewer" dc:"句子审查者 ID"`
	CommitFrom string                `json:"commit_from" dc:"句子提交来源"`
	Status     consts.HitokotoStatus `json:"status" dc:"句子状态"`
	CreatedAt  string                `json:"created_at" dc:"创建时间"`
}

type HitokotoV1WithPoll struct {
	HitokotoV1
	PollStatus consts.PollStatus `json:"poll_status" dc:"句子投票状态"`
}

type GetHitokotoV1SchemaListInput struct {
	Page     int    `json:"page" v:"min:1" d:"1" dc:"页码"`
	PageSize int    `json:"page_size" v:"min:0|max:1000" d:"30" dc:"每页数量"`
	Order    string `json:"order" v:"in:desc,asc" d:"desc" dc:"排序方式"`
	// 搜索参数
	UUIDs    []string               `json:"uuids" dc:"句子 UUID"`
	Keywords *string                `json:"keyword" dc:"句子内容"`
	Creator  *string                `json:"creator" dc:"句子创建者（或 ID）"`
	From     *string                `json:"from" dc:"句子来源"`
	FromWho  *string                `json:"from_who" dc:"句子来源者"`
	Type     *string                `json:"type" dc:"句子类型"`
	Status   *consts.HitokotoStatus `json:"status" dc:"句子状态" v:"in:pending,approved,rejected"`
	Start    *gtime.Time            `json:"start" dc:"开始时间"`
	End      *gtime.Time            `json:"end" dc:"结束时间"`
}

type GetHitokotoV1SchemaListOutput Page[HitokotoV1WithPoll]

type DoHitokotoV1Update struct {
	Hitokoto   *string `json:"hitokoto" dc:"句子内容" v:"length:1,255" mapstructure:"hitokoto,omitempty"`
	From       *string `json:"from" dc:"句子来源" v:"min-length:1" mapstructure:"from,omitempty"`
	FromWho    *string `json:"from_who" dc:"句子来源者" v:"length:1,255" mapstructure:"from_who,omitempty" `
	Type       *string `json:"type" dc:"句子类型" v:"length:1,255" mapstructure:"type,omitempty"`
	CommitFrom *string `json:"commit_from" dc:"句子提交来源" v:"length:1,255" mapstructure:"commit_from,omitempty"`
	Creator    *string `json:"creator" dc:"句子创建者" v:"length:1,255" mapstructure:"creator,omitempty"`
	CreatorUID *uint   `json:"creator_uid" dc:"句子创建者 ID" v:"min:1" mapstructure:"creator_uid,omitempty"`
	Reviewer   *uint   `json:"reviewer" dc:"句子审查者 ID" v:"min:1" mapstructure:"reviewer,omitempty"`
	CreatedAt  *string `json:"created_at" dc:"创建时间" v:"length:1,255" mapstructure:"created_at,omitempty"`
}
