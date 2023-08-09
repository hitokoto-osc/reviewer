// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PollScorePipelineDao is the data access object for table hitokoto_poll_score_pipeline.
type PollScorePipelineDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns PollScorePipelineColumns // columns contains all the column names of Table for convenient usage.
}

// PollScorePipelineColumns defines and stores column names for table hitokoto_poll_score_pipeline.
type PollScorePipelineColumns struct {
	Id           string //
	PollId       string // 触发奖惩的投票
	UserId       string // 操作的用户
	SentenceUuid string // 触发奖惩的句子
	Type         string // 奖励还是惩罚
	Score        string // 分数
	CreatedAt    string //
	UpdatedAt    string //
}

// pollScorePipelineColumns holds the columns for table hitokoto_poll_score_pipeline.
var pollScorePipelineColumns = PollScorePipelineColumns{
	Id:           "id",
	PollId:       "poll_id",
	UserId:       "user_id",
	SentenceUuid: "sentence_uuid",
	Type:         "type",
	Score:        "score",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewPollScorePipelineDao creates and returns a new DAO object for table data access.
func NewPollScorePipelineDao() *PollScorePipelineDao {
	return &PollScorePipelineDao{
		group:   "default",
		table:   "hitokoto_poll_score_pipeline",
		columns: pollScorePipelineColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PollScorePipelineDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PollScorePipelineDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PollScorePipelineDao) Columns() PollScorePipelineColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PollScorePipelineDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PollScorePipelineDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PollScorePipelineDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
