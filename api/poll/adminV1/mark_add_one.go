package adminV1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddOneMarkReq struct {
	g.Meta   `path:"poll/mark" method:"POST" summary:"新建投票标记" tags:"投票标记"`
	Text     string `json:"text"         ` // 标签名称
	Level    string `json:"level"        ` // 严重程度
	Property int    `json:"property"     ` // 分类属性
}

type AddOneMarkRes struct{}
