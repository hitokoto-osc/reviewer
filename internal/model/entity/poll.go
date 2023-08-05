// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Poll is the golang structure for table poll.
type Poll struct {
	Id             int         `json:"id"             ` //
	SentenceUuid   string      `json:"sentenceUuid"   ` // 句子 UUID；标识状态
	Status         int         `json:"status"         ` // 投票状态
	IsExpandedPoll int         `json:"isExpandedPoll" ` // 是否启动了补充投票
	Accept         int         `json:"accept"         ` // 赞成票
	Reject         int         `json:"reject"         ` // 否决票
	NeedEdited     int         `json:"needEdited"     ` // 需要修改
	NeedUserPoll   int         `json:"needUserPoll"   ` //
	PendingId      int         `json:"pendingId"      ` //
	CreatedAt      *gtime.Time `json:"createdAt"      ` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` //
}
