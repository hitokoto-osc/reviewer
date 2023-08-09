// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SentenceHistory is the golang structure for table sentence_history.
type SentenceHistory struct {
	Id           int         `json:"id"           ` //
	SentenceUuid string      `json:"sentenceUuid" ` //
	Before       string      `json:"before"       ` //
	After        string      `json:"after"        ` //
	ModifiedBy   int         `json:"modifiedBy"   ` //
	Reason       string      `json:"reason"       ` //
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
}
