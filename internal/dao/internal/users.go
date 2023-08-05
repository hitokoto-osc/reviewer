// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for table hitokoto_users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for table hitokoto_users.
type UsersColumns struct {
	Id              string //
	Name            string //
	Email           string //
	Token           string //
	Password        string //
	IsSuspended     string //
	IsAdmin         string // 是否是管理员，1为是0为否
	IsReviewer      string // 是否审核者，是：1，不是：0
	RememberToken   string //
	EmailVerifiedAt string // 邮箱验证
	CreatedFrom     string //
	CreatedAt       string //
	UpdatedAt       string //
}

// usersColumns holds the columns for table hitokoto_users.
var usersColumns = UsersColumns{
	Id:              "id",
	Name:            "name",
	Email:           "email",
	Token:           "token",
	Password:        "password",
	IsSuspended:     "is_suspended",
	IsAdmin:         "is_admin",
	IsReviewer:      "is_reviewer",
	RememberToken:   "remember_token",
	EmailVerifiedAt: "email_verified_at",
	CreatedFrom:     "created_from",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "hitokoto_users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
