package adminV1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UpdateOneMarkReq struct {
	g.Meta `path:"poll/mark/:id" method:"PUT" summary:"更新投票标记" tags:"投票标记"`
	// 投票标记 ID
	ID           uint        `json:"id" v:"required#投票标记 ID 不能为空" copier:"Id,must"`
	Text         string      `json:"text"         `  // 标签名称
	Level        string      `json:"level"        `  // 严重程度
	Property     int         `json:"property"     `  // 分类属性
	DeprecatedAt *gtime.Time `json:"deprecated_at" ` //
	UpdatedAt    *gtime.Time `json:"updated_at"    ` // 更新时间
	CreatedAt    *gtime.Time `json:"created_at"    ` //
}

type UpdateOneMarkRes struct{}
