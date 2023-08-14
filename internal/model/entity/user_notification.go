// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserNotification is the golang structure for table user_notification.
type UserNotification struct {
	Id                                int         `json:"id"                                ` //
	UserId                            int         `json:"userId"                            ` //
	EmailNotificationGlobal           int         `json:"emailNotificationGlobal"           ` //
	EmailNotificationHitokotoAppended int         `json:"emailNotificationHitokotoAppended" ` //
	EmailNotificationHitokotoReviewed int         `json:"emailNotificationHitokotoReviewed" ` //
	EmailNotificationPollCreated      int         `json:"emailNotificationPollCreated"      ` //
	EmailNotificationPollResult       int         `json:"emailNotificationPollResult"       ` //
	EmailNotificationPollDailyReport  int         `json:"emailNotificationPollDailyReport"  ` //
	UpdatedAt                         *gtime.Time `json:"updatedAt"                         ` //
	CreatedAt                         *gtime.Time `json:"createdAt"                         ` //
}
