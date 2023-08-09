// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PollLog is the golang structure for table poll_log.
type PollLog struct {
	Id           int         `json:"id"           ` //
	PollId       int         `json:"pollId"       ` // 投票 ID
	UserId       int         `json:"userId"       ` //
	Point        int         `json:"point"        ` // 票数
	SentenceUuid string      `json:"sentenceUuid" ` // 句子 UUID
	Type         int         `json:"type"         ` // 投票种类
	Comment      string      `json:"comment"      ` //
	IsAdmin      int         `json:"isAdmin"      ` // 是否是管理，追踪一票否决。
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
}
