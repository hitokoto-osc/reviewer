package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetUserPollLogReq struct {
	g.Meta `path:"/user/poll/log" tags:"User" method:"get" summary:"获取用户投票记录"`
	Offset int `json:"offset" dc:"偏移量" d:"0" v:"integer|min-length:0#偏移量必须为整数|偏移量不能小于0"`
	Limit  int `json:"limit" dc:"数量" d:"30" v:"integer|length:0,1000#数量必须为整数|数量不能超过1000"`
}

type GetUserPollLogRes []model.UserPollLog
