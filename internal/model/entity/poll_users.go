// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PollUsers is the golang structure for table poll_users.
type PollUsers struct {
	Id           int         `json:"id"           ` //
	UserId       int         `json:"userId"       ` //
	Points       int         `json:"points"       ` // 总票数
	Accept       int         `json:"accept"       ` // 赞同票
	Reject       int         `json:"reject"       ` // 拒绝票
	NeedEdited   int         `json:"needEdited"   ` // 需要修改
	Score        int         `json:"score"        ` // 分数
	AdoptionRate float64     `json:"adoptionRate" ` // 采纳率
	CreatedAt    *gtime.Time `json:"createdAt"    ` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` //
}
