package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetUserScoreRecordsReq struct {
	g.Meta   `path:"/user/score/records" tags:"User" method:"get" summary:"获得用户积分记录"`
	Page     int    `json:"page" dc:"页码" d:"1" v:"integer|min-length:1#页码必须为整数|页码不能小于1"`
	PageSize int    `json:"page_size" dc:"页面大小" d:"30" v:"integer|length:0,1000#页面大小必须为整数|页面大小不能超过1000"`
	Order    string `json:"order" dc:"排序" d:"desc" v:"in:desc,asc#排序方式不正确"`
}

type GetUserScoreRecordsRes struct {
	model.GetUserScoreRecordsOutput
}
