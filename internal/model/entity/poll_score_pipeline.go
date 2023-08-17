// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PollScorePipeline is the golang structure for table poll_score_pipeline.
type PollScorePipeline struct {
	Id           int         `json:"id"           ` //
	PollId       int         `json:"pollId"       ` // 触发奖惩的投票
	UserId       int         `json:"userId"       ` // 操作的用户
	SentenceUuid string      `json:"sentenceUuid" ` // 触发奖惩的句子
	Type         int         `json:"type"         ` // 奖励还是惩罚
	Score        int         `json:"score"        ` // 分数
	Reason       string      `json:"reason"       ` // 理由
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
}
