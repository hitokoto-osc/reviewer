// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PollMark is the golang structure for table poll_mark.
type PollMark struct {
	Id        uint        `json:"id"        ` //
	Text      string      `json:"text"      ` // 标签名称
	Level     string      `json:"level"     ` // 严重程度
	Property  int         `json:"property"  ` // 分类属性
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
	CreatedAt *gtime.Time `json:"createdAt" ` //
}
