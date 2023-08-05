// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PollLogDao is the data access object for table hitokoto_poll_log.
type PollLogDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns PollLogColumns // columns contains all the column names of Table for convenient usage.
}

// PollLogColumns defines and stores column names for table hitokoto_poll_log.
type PollLogColumns struct {
	Id           string //
	UserId       string //
	Point        string // 票数
	SentenceUuid string // 句子 UUID
	Type         string // 投票种类
	Comment      string //
	IsAdmin      string // 是否是管理，追踪一票否决。
	CreatedAt    string //
	UpdatedAt    string //
}

// pollLogColumns holds the columns for table hitokoto_poll_log.
var pollLogColumns = PollLogColumns{
	Id:           "id",
	UserId:       "user_id",
	Point:        "point",
	SentenceUuid: "sentence_uuid",
	Type:         "type",
	Comment:      "comment",
	IsAdmin:      "is_admin",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewPollLogDao creates and returns a new DAO object for table data access.
func NewPollLogDao() *PollLogDao {
	return &PollLogDao{
		group:   "default",
		table:   "hitokoto_poll_log",
		columns: pollLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PollLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PollLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PollLogDao) Columns() PollLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PollLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PollLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PollLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
