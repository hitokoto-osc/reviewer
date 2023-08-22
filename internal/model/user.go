package model

import (
	"github.com/gogf/gf/v2/database/gdb"
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
	Points       UserPollPoints `json:"points" dc:"投票点数"`
	Count        int            `json:"count" dc:"投票次数"`
	Score        int            `json:"score" dc:"积分"`
	AdoptionRate float64        `json:"adoption_rate" dc:"采纳率"`
	CreatedAt    *time.Time     `json:"created_at" dc:"创建时间"`
	UpdatedAt    *time.Time     `json:"updated_at" dc:"更新时间"`
}

type UserPollLog struct {
	PollID       uint              `json:"poll_id" dc:"投票 ID"`
	Point        int               `json:"point" dc:"投票点数"`
	SentenceUUID string            `json:"sentence_uuid" dc:"句子 UUID"`
	Method       consts.PollMethod `json:"type" dc:"投票类型"`
	Comment      string            `json:"comment" dc:"理由"`
	CreatedAt    *time.Time        `json:"created_at" dc:"投票时间"`
	UpdatedAt    *time.Time        `json:"updated_at" dc:"更新时间"`
}

type UserPollLogWithSentence struct {
	UserPollLog
	Sentence *HitokotoV1Schema `json:"sentence" dc:"句子信息"`
}

type UserPollElement struct {
	UserPollLogWithSentence
	PollInfo *PollElement `json:"poll_info" dc:"投票信息"`
	Marks    []int        `json:"marks" dc:"投票标记"`
}

type UserPollResult struct {
	Total uint `json:"total" dc:"总数"`
	GetUserPollLogsWithPollResultOutput
}

type GetUserPollLogsInput struct {
	UserID    uint   // 用户 ID
	Order     string // 排序方式 `dc:"排序方式"`
	Page      int
	PageSize  int
	WithCache bool
}

type GetUserPollLogsOutput struct {
	Collection []UserPollLog `json:"collection" dc:"数据"`
	Page       int           `json:"page" dc:"当前页数"`
	PageSize   int           `json:"page_size" dc:"每页数量"`
}

type GetUserPollLogsWithSentenceInput = GetUserPollLogsInput

type GetUserPollLogsWithSentenceOutput struct {
	Collection []UserPollLogWithSentence `json:"collection" dc:"数据"`
	Page       int                       `json:"page" dc:"当前页数"`
	PageSize   int                       `json:"page_size" dc:"每页数量"`
}

type GetUserPollLogsWithPollResultInput = GetUserPollLogsInput

type GetUserPollLogsWithPollResultOutput struct {
	Collection []UserPollElement `json:"collection" dc:"数据"`
	Page       int               `json:"page" dc:"当前页数"`
	PageSize   int               `json:"page_size" dc:"每页数量"`
}

type UserPollScoreInput struct {
	UserID       uint
	PollID       uint
	Score        int
	SentenceUUID string
	Reason       string
	Tx           gdb.TX
}
