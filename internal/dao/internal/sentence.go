// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SentenceDao is the data access object for table hitokoto_sentence.
type SentenceDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns SentenceColumns // columns contains all the column names of Table for convenient usage.
}

// SentenceColumns defines and stores column names for table hitokoto_sentence.
type SentenceColumns struct {
	Id         string //
	Uuid       string //
	Hitokoto   string //
	Type       string //
	From       string //
	FromWho    string //
	Creator    string //
	CreatorUid string //
	Reviewer   string //
	CommitFrom string //
	Assessor   string //
	Owner      string //
	CreatedAt  string //
}

// sentenceColumns holds the columns for table hitokoto_sentence.
var sentenceColumns = SentenceColumns{
	Id:         "id",
	Uuid:       "uuid",
	Hitokoto:   "hitokoto",
	Type:       "type",
	From:       "from",
	FromWho:    "from_who",
	Creator:    "creator",
	CreatorUid: "creator_uid",
	Reviewer:   "reviewer",
	CommitFrom: "commit_from",
	Assessor:   "assessor",
	Owner:      "owner",
	CreatedAt:  "created_at",
}

// NewSentenceDao creates and returns a new DAO object for table data access.
func NewSentenceDao() *SentenceDao {
	return &SentenceDao{
		group:   "default",
		table:   "hitokoto_sentence",
		columns: sentenceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SentenceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SentenceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SentenceDao) Columns() SentenceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SentenceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SentenceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SentenceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
