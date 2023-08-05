// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PollLog is the golang structure of table hitokoto_poll_log for DAO operations like Where/Data.
type PollLog struct {
	g.Meta       `orm:"table:hitokoto_poll_log, do:true"`
	Id           interface{} //
	UserId       interface{} //
	Point        interface{} // 票数
	SentenceUuid interface{} // 句子 UUID
	Type         interface{} // 投票种类
	Comment      interface{} //
	IsAdmin      interface{} // 是否是管理，追踪一票否决。
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
