// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SentenceHistory is the golang structure of table hitokoto_sentence_history for DAO operations like Where/Data.
type SentenceHistory struct {
	g.Meta       `orm:"table:hitokoto_sentence_history, do:true"`
	Id           interface{} //
	SentenceUuid interface{} //
	Before       interface{} //
	After        interface{} //
	ModifiedBy   interface{} //
	Reason       interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
