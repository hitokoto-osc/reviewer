package model

import (
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/utility/time"
)

type GetUserScoreRecordsInput struct {
	UserID    uint   `json:"user_id" dc:"用户 ID"`
	Page      int    `json:"page" dc:"当前页数"`
	PageSize  int    `json:"page_size" dc:"每页数量"`
	Order     string `json:"order" dc:"排序方式"`
	WithCache bool
}

type UserScoreRecord struct {
	ID           uint                 `json:"id" dc:"ID"`
	PollID       uint                 `json:"poll_id" dc:"投票 ID"`
	UserID       uint                 `json:"user_id" dc:"用户 ID"`
	SentenceUUID string               `json:"sentence_uuid" dc:"句子 UUID"`
	Score        int                  `json:"score" dc:"积分"`
	Type         consts.UserScoreType `json:"type" dc:"类型"`
	Reason       string               `json:"reason" dc:"原因"`
	UpdatedAt    *time.Time           `json:"updated_at" dc:"更新时间"`
	CreatedAt    *time.Time           `json:"created_at" dc:"创建时间"`
}

type GetUserScoreRecordsOutput struct {
	Collection []UserScoreRecord `json:"collection" dc:"数据"`
	Total      int               `json:"total" dc:"总数"`
	Page       int               `json:"page" dc:"当前页数"`
	PageSize   int               `json:"page_size" dc:"每页数量"`
}
