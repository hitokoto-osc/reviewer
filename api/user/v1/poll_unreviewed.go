package v1

import "github.com/gogf/gf/v2/frame/g"

type GetUserPollUnreviewedReq struct {
	g.Meta `path:"/user/poll/unreviewed" tags:"User" method:"get" summary:"获得未审查的投票数量"`
}

type GetUserPollUnreviewedRes struct {
	Count int `json:"count" dc:"未审查的投票数量"`
}
