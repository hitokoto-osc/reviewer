// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PollMarkRelation is the golang structure for table poll_mark_relation.
type PollMarkRelation struct {
	Id           int         `json:"id"           ` //
	UserId       int         `json:"userId"       ` //
	SentenceUuid string      `json:"sentenceUuid" ` //
	IsRefuse     uint        `json:"isRefuse"     ` //
	MarkId       int         `json:"markId"       ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
}
