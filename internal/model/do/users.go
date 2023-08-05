// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table hitokoto_users for DAO operations like Where/Data.
type Users struct {
	g.Meta          `orm:"table:hitokoto_users, do:true"`
	Id              interface{} //
	Name            interface{} //
	Email           interface{} //
	Token           interface{} //
	Password        interface{} //
	IsSuspended     interface{} //
	IsAdmin         interface{} // 是否是管理员，1为是0为否
	IsReviewer      interface{} // 是否审核者，是：1，不是：0
	RememberToken   interface{} //
	EmailVerifiedAt *gtime.Time // 邮箱验证
	CreatedFrom     interface{} //
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
