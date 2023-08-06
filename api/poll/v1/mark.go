package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type GetPollMarksReq struct {
	g.Meta `path:"/poll/mark" tags:"Poll" method:"get" summary:"获取投票标记"`
}

type GetPollMarksRes []model.PollMark
