package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type GetPollMarksReq struct {
	g.Meta `path:"/poll/mark" tags:"Poll" method:"get" summary:"获取投票标记"`
}

type PollMark struct {
	ID        int64                   `json:"id" dc:"标记 ID"`
	Text      string                  `json:"text" dc:"标记文本"`
	Level     consts.PollMarkLevel    `json:"level" dc:"标记等级"`
	Property  consts.PollMarkProperty `json:"property" dc:"标记属性"`
	UpdatedAt string                  `json:"updated_at" dc:"更新时间"`
	CreatedAt string                  `json:"created_at" dc:"创建时间"`
}
type GetPollMarksRes []PollMark
