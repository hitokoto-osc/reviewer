// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PollScorePipeline is the golang structure of table hitokoto_poll_score_pipeline for DAO operations like Where/Data.
type PollScorePipeline struct {
	g.Meta       `orm:"table:hitokoto_poll_score_pipeline, do:true"`
	Id           interface{} //
	UserId       interface{} // 操作的用户
	SentenceUuid interface{} // 触发奖惩的句子
	Type         interface{} // 奖励还是惩罚
	Score        interface{} // 分数
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
