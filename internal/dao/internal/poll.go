// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PollDao is the data access object for table hitokoto_poll.
type PollDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns PollColumns // columns contains all the column names of Table for convenient usage.
}

// PollColumns defines and stores column names for table hitokoto_poll.
type PollColumns struct {
	Id             string //
	SentenceUuid   string // 句子 UUID；标识状态
	Status         string // 投票状态
	IsExpandedPoll string // 是否启动了补充投票
	Accept         string // 赞成票
	Reject         string // 否决票
	NeedEdited     string // 需要修改
	NeedUserPoll   string //
	PendingId      string //
	CreatedAt      string //
	UpdatedAt      string //
}

// pollColumns holds the columns for table hitokoto_poll.
var pollColumns = PollColumns{
	Id:             "id",
	SentenceUuid:   "sentence_uuid",
	Status:         "status",
	IsExpandedPoll: "is_expanded_poll",
	Accept:         "accept",
	Reject:         "reject",
	NeedEdited:     "need_edited",
	NeedUserPoll:   "need_user_poll",
	PendingId:      "pending_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewPollDao creates and returns a new DAO object for table data access.
func NewPollDao() *PollDao {
	return &PollDao{
		group:   "default",
		table:   "hitokoto_poll",
		columns: pollColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PollDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PollDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PollDao) Columns() PollColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PollDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PollDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PollDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
