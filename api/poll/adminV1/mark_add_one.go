package adminV1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddOneMarkReq struct {
	g.Meta `path:"poll/mark" method:"POST" summary:"新建投票标记" tags:"投票标记"`
	// 投票标记 ID
	ID           uint        `json:"id" v:"required#投票标记 ID 不能为空" copier:"Id,must"`
	Text         string      `json:"text"         `  // 标签名称
	Level        string      `json:"level"        `  // 严重程度
	Property     int         `json:"property"     `  // 分类属性
	DeprecatedAt *gtime.Time `json:"deprecated_at" ` //
}

type AddOneMarkRes struct{}
