package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

type GetUserPollResultReq struct {
	g.Meta `path:"/user/poll/result" tags:"User" method:"get" summary:"获取用户投票结果"`
	Offset int `json:"offset" dc:"偏移量" v:"required|integer#偏移量必须为整数"`
	Limit  int `json:"limit" dc:"数量" v:"required|integer#数量必须为整数"`
}

type UserPollElement struct {
	UserPollLog
	PollInfo v1.PollSchema `json:"poll_info" dc:"投票信息"`
	Marks    []int64       `json:"marks" dc:"投票标记"`
}

type UserPollResult struct {
	Total      int64             `json:"total" dc:"总数"`
	Collection []UserPollElement `json:"collection" dc:"数据"`
}

type GetUserPollResultRes []UserPollResult
