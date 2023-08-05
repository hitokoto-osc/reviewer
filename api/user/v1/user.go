package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type GetUserReq struct {
	g.Meta       `path:"/user" tags:"User" method:"get" summary:"获得用户信息"`
	WithPollLogs bool `json:"with_poll_logs" dc:"是否返回投票记录" v:"boolean" d:"false"`
}

type UserPollPoints struct {
	Total      int `json:"total" dc:"总投票点数"`
	Approve    int `json:"approved" dc:"赞成票"`
	Reject     int `json:"rejected" dc:"反对票"`
	NeedModify int `json:"need_modify" dc:"需修改票"`
}

type UserPoll struct {
	Points    UserPollPoints `json:"points" dc:"投票点数"`
	Count     int            `json:"count" dc:"投票次数"`
	Score     int            `json:"score" dc:"投票得分"`
	CreatedAt string         `json:"created_at" dc:"创建时间"`
	UpdatedAt string         `json:"updated_at" dc:"更新时间"`
}

type GetUserRes struct {
	ID        int64           `json:"id" dc:"用户 ID"`
	Name      string          `json:"name" dc:"用户名"`
	Email     string          `json:"email" dc:"邮箱"`
	Role      consts.UserRole `json:"role" dc:"角色"`
	Poll      UserPoll        `json:"poll" dc:"投票信息"`
	CreatedAt string          `json:"created_at" dc:"创建时间"`
	UpdatedAt string          `json:"updated_at" dc:"更新时间"`
}
