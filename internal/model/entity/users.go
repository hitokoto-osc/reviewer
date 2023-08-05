// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id              uint        `json:"id"              ` //
	Name            string      `json:"name"            ` //
	Email           string      `json:"email"           ` //
	Token           string      `json:"token"           ` //
	Password        string      `json:"password"        ` //
	IsSuspended     int         `json:"isSuspended"     ` //
	IsAdmin         uint        `json:"isAdmin"         ` // 是否是管理员，1为是0为否
	IsReviewer      int         `json:"isReviewer"      ` // 是否审核者，是：1，不是：0
	RememberToken   string      `json:"rememberToken"   ` //
	EmailVerifiedAt *gtime.Time `json:"emailVerifiedAt" ` // 邮箱验证
	CreatedFrom     string      `json:"createdFrom"     ` //
	CreatedAt       *gtime.Time `json:"createdAt"       ` //
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` //
}
