// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PollPipeline is the golang structure for table poll_pipeline.
type PollPipeline struct {
	Id           int         `json:"id"           ` //
	SentenceUuid string      `json:"sentenceUuid" ` //
	Operate      int         `json:"operate"      ` //
	Mark         string      `json:"mark"         ` //
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
}
