package adminV1

import "github.com/gogf/gf/v2/frame/g"

type DeleteListReq struct {
	g.Meta `path:"/hitokoto" method:"DELETE" summary:"批量删除句子" tags:"句子"`
	UUIDs  []string `json:"uuids" dc:"句子 UUID" v:"required|length:1,100" in:"query"`
}

type DeleteListRes struct{}
