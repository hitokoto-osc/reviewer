// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserNotificationDao is the data access object for table hitokoto_user_notification.
type UserNotificationDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UserNotificationColumns // columns contains all the column names of Table for convenient usage.
}

// UserNotificationColumns defines and stores column names for table hitokoto_user_notification.
type UserNotificationColumns struct {
	Id                                string //
	UserId                            string //
	EmailNotificationGlobal           string //
	EmailNotificationHitokotoAppended string //
	EmailNotificationHitokotoReviewed string //
	EmailNotificationPollCreated      string //
	EmailNotificationPollResult       string //
	EmailNotificationPollDailyReport  string //
	UpdatedAt                         string //
	CreatedAt                         string //
}

// userNotificationColumns holds the columns for table hitokoto_user_notification.
var userNotificationColumns = UserNotificationColumns{
	Id:                                "id",
	UserId:                            "user_id",
	EmailNotificationGlobal:           "email_notification_global",
	EmailNotificationHitokotoAppended: "email_notification_hitokoto_appended",
	EmailNotificationHitokotoReviewed: "email_notification_hitokoto_reviewed",
	EmailNotificationPollCreated:      "email_notification_poll_created",
	EmailNotificationPollResult:       "email_notification_poll_result",
	EmailNotificationPollDailyReport:  "email_notification_poll_daily_report",
	UpdatedAt:                         "updated_at",
	CreatedAt:                         "created_at",
}

// NewUserNotificationDao creates and returns a new DAO object for table data access.
func NewUserNotificationDao() *UserNotificationDao {
	return &UserNotificationDao{
		group:   "default",
		table:   "hitokoto_user_notification",
		columns: userNotificationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserNotificationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserNotificationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserNotificationDao) Columns() UserNotificationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserNotificationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserNotificationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserNotificationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
