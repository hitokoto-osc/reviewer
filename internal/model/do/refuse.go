// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Refuse is the golang structure of table hitokoto_refuse for DAO operations like Where/Data.
type Refuse struct {
	g.Meta     `orm:"table:hitokoto_refuse, do:true"`
	Id         interface{} //
	Uuid       interface{} //
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
