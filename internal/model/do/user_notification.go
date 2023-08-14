// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserNotification is the golang structure of table hitokoto_user_notification for DAO operations like Where/Data.
type UserNotification struct {
	g.Meta                            `orm:"table:hitokoto_user_notification, do:true"`
	Id                                interface{} //
	UserId                            interface{} //
	EmailNotificationGlobal           interface{} //
	EmailNotificationHitokotoAppended interface{} //
	EmailNotificationHitokotoReviewed interface{} //
	EmailNotificationPollCreated      interface{} //
	EmailNotificationPollResult       interface{} //
	EmailNotificationPollDailyReport  interface{} //
	UpdatedAt                         *gtime.Time //
	CreatedAt                         *gtime.Time //
}
