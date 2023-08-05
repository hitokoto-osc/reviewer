// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PollPipeline is the golang structure of table hitokoto_poll_pipeline for DAO operations like Where/Data.
type PollPipeline struct {
	g.Meta       `orm:"table:hitokoto_poll_pipeline, do:true"`
	Id           interface{} //
	SentenceUuid interface{} //
	Operate      interface{} //
	Mark         interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
