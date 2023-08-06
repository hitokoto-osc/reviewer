package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetUserPollLogReq struct {
	g.Meta `path:"/user/poll/log" tags:"User" method:"get" summary:"获取用户投票记录"`
	Offset int `json:"offset" dc:"偏移量" v:"required|integer#偏移量必须为整数"`
	Limit  int `json:"limit" dc:"数量" v:"required|integer#数量必须为整数"`
}

type GetUserPollLogRes []model.UserPollLog
