package model

import (
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type UserPattern struct {
	entity.Users
	Status consts.UserStatus `json:"status" dc:"用户状态"`
	Role   consts.UserRole   `json:"role" dc:"用户角色"`
	Poll   entity.PollUsers  `json:"poll" dc:"用户投票信息"`
}
