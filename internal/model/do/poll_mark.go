// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PollMark is the golang structure of table hitokoto_poll_mark for DAO operations like Where/Data.
type PollMark struct {
	g.Meta    `orm:"table:hitokoto_poll_mark, do:true"`
	Id        interface{} //
	Text      interface{} // 标签名称
	Level     interface{} // 严重程度
	Property  interface{} // 分类属性
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time //
}
