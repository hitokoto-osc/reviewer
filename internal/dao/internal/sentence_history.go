// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SentenceHistoryDao is the data access object for table hitokoto_sentence_history.
type SentenceHistoryDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns SentenceHistoryColumns // columns contains all the column names of Table for convenient usage.
}

// SentenceHistoryColumns defines and stores column names for table hitokoto_sentence_history.
type SentenceHistoryColumns struct {
	Id           string //
	SentenceUuid string //
	Before       string //
	After        string //
	ModifiedBy   string //
	Reason       string //
	CreatedAt    string //
	UpdatedAt    string //
}

// sentenceHistoryColumns holds the columns for table hitokoto_sentence_history.
var sentenceHistoryColumns = SentenceHistoryColumns{
	Id:           "id",
	SentenceUuid: "sentence_uuid",
	Before:       "before",
	After:        "after",
	ModifiedBy:   "modified_by",
	Reason:       "reason",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSentenceHistoryDao creates and returns a new DAO object for table data access.
func NewSentenceHistoryDao() *SentenceHistoryDao {
	return &SentenceHistoryDao{
		group:   "default",
		table:   "hitokoto_sentence_history",
		columns: sentenceHistoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SentenceHistoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SentenceHistoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SentenceHistoryDao) Columns() SentenceHistoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SentenceHistoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SentenceHistoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SentenceHistoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
