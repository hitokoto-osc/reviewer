// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Sentence is the golang structure of table hitokoto_sentence for DAO operations like Where/Data.
type Sentence struct {
	g.Meta     `orm:"table:hitokoto_sentence, do:true"`
	Id         interface{} //
	Uuid       interface{} //
	Hitokoto   interface{} //
	Type       interface{} //
	From       interface{} //
	FromWho    interface{} //
	Creator    interface{} //
	CreatorUid interface{} //
	Reviewer   interface{} //
	CommitFrom interface{} //
	Assessor   interface{} //
	Owner      interface{} //
	CreatedAt  interface{} //
}
