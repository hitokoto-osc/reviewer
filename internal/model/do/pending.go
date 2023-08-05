// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Pending is the golang structure of table hitokoto_pending for DAO operations like Where/Data.
type Pending struct {
	g.Meta     `orm:"table:hitokoto_pending, do:true"`
	Id         interface{} //
	Uuid       interface{} //
	PollStatus interface{} // 投票状态，默认： 0，未启用投票
	Hitokoto   interface{} //
	Type       interface{} //
	From       interface{} //
	FromWho    interface{} //
	Creator    interface{} //
	CreatorUid interface{} //
	Owner      interface{} //
	Reviewer   interface{} //
	CommitFrom interface{} //
	CreatedAt  interface{} //
}
