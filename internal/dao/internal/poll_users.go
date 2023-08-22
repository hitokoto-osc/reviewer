// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PollUsersDao is the data access object for table hitokoto_poll_users.
type PollUsersDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns PollUsersColumns // columns contains all the column names of Table for convenient usage.
}

// PollUsersColumns defines and stores column names for table hitokoto_poll_users.
type PollUsersColumns struct {
	Id           string //
	UserId       string //
	Points       string // 总票数
	Accept       string // 赞同票
	Reject       string // 拒绝票
	NeedEdited   string // 需要修改
	Score        string // 分数
	AdoptionRate string // 采纳率
	CreatedAt    string //
	UpdatedAt    string //
}

// pollUsersColumns holds the columns for table hitokoto_poll_users.
var pollUsersColumns = PollUsersColumns{
	Id:           "id",
	UserId:       "user_id",
	Points:       "points",
	Accept:       "accept",
	Reject:       "reject",
	NeedEdited:   "need_edited",
	Score:        "score",
	AdoptionRate: "adoption_rate",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewPollUsersDao creates and returns a new DAO object for table data access.
func NewPollUsersDao() *PollUsersDao {
	return &PollUsersDao{
		group:   "default",
		table:   "hitokoto_poll_users",
		columns: pollUsersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PollUsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PollUsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PollUsersDao) Columns() PollUsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PollUsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PollUsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PollUsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
