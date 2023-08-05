// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PollMarkRelation is the golang structure of table hitokoto_poll_mark_relation for DAO operations like Where/Data.
type PollMarkRelation struct {
	g.Meta       `orm:"table:hitokoto_poll_mark_relation, do:true"`
	Id           interface{} //
	UserId       interface{} //
	SentenceUuid interface{} //
	IsRefuse     interface{} //
	MarkId       interface{} //
	UpdatedAt    *gtime.Time //
	CreatedAt    *gtime.Time //
}
