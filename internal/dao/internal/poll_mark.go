// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PollMarkDao is the data access object for table hitokoto_poll_mark.
type PollMarkDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns PollMarkColumns // columns contains all the column names of Table for convenient usage.
}

// PollMarkColumns defines and stores column names for table hitokoto_poll_mark.
type PollMarkColumns struct {
	Id        string //
	Text      string // 标签名称
	Level     string // 严重程度
	Property  string // 分类属性
	UpdatedAt string // 更新时间
	CreatedAt string //
}

// pollMarkColumns holds the columns for table hitokoto_poll_mark.
var pollMarkColumns = PollMarkColumns{
	Id:        "id",
	Text:      "text",
	Level:     "level",
	Property:  "property",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// NewPollMarkDao creates and returns a new DAO object for table data access.
func NewPollMarkDao() *PollMarkDao {
	return &PollMarkDao{
		group:   "default",
		table:   "hitokoto_poll_mark",
		columns: pollMarkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PollMarkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PollMarkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PollMarkDao) Columns() PollMarkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PollMarkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PollMarkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PollMarkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
