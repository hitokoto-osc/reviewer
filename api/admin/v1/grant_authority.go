package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type GrantUserAuthorityReq struct {
	g.Meta      `path:"/admin/grant_authority" tags:"Admin" method:"post" summary:"授予用户权限"`
	Account     string               `json:"account" dc:"用户账号" v:"required"`
	AccountType consts.UserLoginType `json:"account_type" dc:"用户账号类型" v:"required|enum#uid,email,token,username"`
}

type GrantUserAuthorityRes struct {
	ID        uint              `json:"id" dc:"用户 ID"`
	Name      string            `json:"name" dc:"用户名"`
	Email     string            `json:"email" dc:"邮箱"`
	Token     string            `json:"token" dc:"Token"`
	Status    consts.UserStatus `json:"status" dc:"用户状态"`
	Role      consts.UserRole   `json:"role" dc:"角色"`
	CreatedAt string            `json:"created_at" dc:"创建时间"`
	UpdatedAt string            `json:"updated_at" dc:"更新时间"`
}
