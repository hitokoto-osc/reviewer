// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PollUsers is the golang structure of table hitokoto_poll_users for DAO operations like Where/Data.
type PollUsers struct {
	g.Meta       `orm:"table:hitokoto_poll_users, do:true"`
	Id           interface{} //
	UserId       interface{} //
	Points       interface{} // 总票数
	Accept       interface{} // 赞同票
	Reject       interface{} // 拒绝票
	NeedEdited   interface{} // 需要修改
	Score        interface{} // 分数
	AdoptionRate interface{} // 采纳率
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
