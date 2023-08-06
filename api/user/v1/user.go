package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/utility/time"
)

type GetUserReq struct {
	g.Meta       `path:"/user" tags:"User" method:"get" summary:"获得用户信息"`
	WithPollLogs bool `json:"with_poll_logs" dc:"是否返回投票记录" v:"boolean" d:"false"`
}

type GetUserRes struct {
	ID        uint                `json:"id" dc:"用户 ID"`
	Name      string              `json:"name" dc:"用户名"`
	Email     string              `json:"email" dc:"邮箱"`
	Role      consts.UserRole     `json:"role" dc:"角色"`
	Poll      model.UserPoll      `json:"poll" dc:"投票信息"`
	PollLog   []model.UserPollLog `json:"poll_log" dc:"投票记录"`
	CreatedAt *time.Time          `json:"created_at" dc:"创建时间"`
	UpdatedAt *time.Time          `json:"updated_at" dc:"更新时间"`
}
