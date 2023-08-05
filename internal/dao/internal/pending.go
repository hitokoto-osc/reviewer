// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PendingDao is the data access object for table hitokoto_pending.
type PendingDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns PendingColumns // columns contains all the column names of Table for convenient usage.
}

// PendingColumns defines and stores column names for table hitokoto_pending.
type PendingColumns struct {
	Id         string //
	Uuid       string //
	PollStatus string // 投票状态，默认： 0，未启用投票
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

// pendingColumns holds the columns for table hitokoto_pending.
var pendingColumns = PendingColumns{
	Id:         "id",
	Uuid:       "uuid",
	PollStatus: "poll_status",
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

// NewPendingDao creates and returns a new DAO object for table data access.
func NewPendingDao() *PendingDao {
	return &PendingDao{
		group:   "default",
		table:   "hitokoto_pending",
		columns: pendingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PendingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PendingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PendingDao) Columns() PendingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PendingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PendingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PendingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
