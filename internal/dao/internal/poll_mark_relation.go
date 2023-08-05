// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PollMarkRelationDao is the data access object for table hitokoto_poll_mark_relation.
type PollMarkRelationDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns PollMarkRelationColumns // columns contains all the column names of Table for convenient usage.
}

// PollMarkRelationColumns defines and stores column names for table hitokoto_poll_mark_relation.
type PollMarkRelationColumns struct {
	Id           string //
	UserId       string //
	SentenceUuid string //
	IsRefuse     string //
	MarkId       string //
	UpdatedAt    string //
	CreatedAt    string //
}

// pollMarkRelationColumns holds the columns for table hitokoto_poll_mark_relation.
var pollMarkRelationColumns = PollMarkRelationColumns{
	Id:           "id",
	UserId:       "user_id",
	SentenceUuid: "sentence_uuid",
	IsRefuse:     "is_refuse",
	MarkId:       "mark_id",
	UpdatedAt:    "updated_at",
	CreatedAt:    "created_at",
}

// NewPollMarkRelationDao creates and returns a new DAO object for table data access.
func NewPollMarkRelationDao() *PollMarkRelationDao {
	return &PollMarkRelationDao{
		group:   "default",
		table:   "hitokoto_poll_mark_relation",
		columns: pollMarkRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PollMarkRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PollMarkRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PollMarkRelationDao) Columns() PollMarkRelationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PollMarkRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PollMarkRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PollMarkRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
