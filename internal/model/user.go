package model

import (
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/utility/time"
)

// UserCtxSchema 用户上下文模型，存储在 bizctx 中
type UserCtxSchema struct {
	entity.Users
	Status consts.UserStatus `json:"status" dc:"用户状态"`
	Role   consts.UserRole   `json:"role" dc:"用户角色"`
	Poll   entity.PollUsers  `json:"poll" dc:"用户投票信息"`
}

type UserPollPoints struct {
	Total      int `json:"total" dc:"总投票点数"`
	Approve    int `json:"approved" dc:"赞成票"`
	Reject     int `json:"rejected" dc:"反对票"`
	NeedModify int `json:"need_modify" dc:"需修改票"`
}

type UserPoll struct {
	Points    UserPollPoints `json:"points" dc:"投票点数"`
	Count     int            `json:"count" dc:"投票次数"`
	Score     int            `json:"score" dc:"投票得分"`
	CreatedAt *time.Time     `json:"created_at" dc:"创建时间"`
	UpdatedAt *time.Time     `json:"updated_at" dc:"更新时间"`
}

type UserPollLog struct {
	Point        int               `json:"point" dc:"投票点数"`
	SentenceUUID string            `json:"sentence_uuid" dc:"句子 UUID"`
	Sentence     *HitokotoV1Schema `json:"sentence" dc:"句子信息"`
	Method       consts.PollMethod `json:"type" dc:"投票类型"`
	Comment      string            `json:"comment" dc:"理由"`
	CreatedAt    *time.Time        `json:"created_at" dc:"投票时间"`
	UpdatedAt    *time.Time        `json:"updated_at" dc:"更新时间"`
}

type UserPollElement struct {
	UserPollLog
	PollInfo PollSchema `json:"poll_info" dc:"投票信息"`
	Marks    []uint     `json:"marks" dc:"投票标记"`
}

type UserPollResult struct {
	Total      uint              `json:"total" dc:"总数"`
	Collection []UserPollElement `json:"collection" dc:"数据"`
}
