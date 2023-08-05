// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RefuseDao is the data access object for table hitokoto_refuse.
type RefuseDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns RefuseColumns // columns contains all the column names of Table for convenient usage.
}

// RefuseColumns defines and stores column names for table hitokoto_refuse.
type RefuseColumns struct {
	Id         string //
	Uuid       string //
	Hitokoto   string //
	Type       string //
	From       string //
	FromWho    string //
	Creator    string //
	CreatorUid string //
	Owner      string //
	Reviewer   string //
	CommitFrom string //
	CreatedAt  string //
}

// refuseColumns holds the columns for table hitokoto_refuse.
var refuseColumns = RefuseColumns{
	Id:         "id",
	Uuid:       "uuid",
	Hitokoto:   "hitokoto",
	Type:       "type",
	From:       "from",
	FromWho:    "from_who",
	Creator:    "creator",
	CreatorUid: "creator_uid",
	Owner:      "owner",
	Reviewer:   "reviewer",
	CommitFrom: "commit_from",
	CreatedAt:  "created_at",
}

// NewRefuseDao creates and returns a new DAO object for table data access.
func NewRefuseDao() *RefuseDao {
	return &RefuseDao{
		group:   "default",
		table:   "hitokoto_refuse",
		columns: refuseColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RefuseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RefuseDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RefuseDao) Columns() RefuseColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RefuseDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RefuseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RefuseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
