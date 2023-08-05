// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Poll is the golang structure of table hitokoto_poll for DAO operations like Where/Data.
type Poll struct {
	g.Meta         `orm:"table:hitokoto_poll, do:true"`
	Id             interface{} //
	SentenceUuid   interface{} // 句子 UUID；标识状态
	Status         interface{} // 投票状态
	IsExpandedPoll interface{} // 是否启动了补充投票
	Accept         interface{} // 赞成票
	Reject         interface{} // 否决票
	NeedEdited     interface{} // 需要修改
	NeedUserPoll   interface{} //
	PendingId      interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
