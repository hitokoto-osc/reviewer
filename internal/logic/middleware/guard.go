package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

// GuardV1 是 v1 用于检查用户权限的中间件
func (s *sMiddleware) GuardV1(role consts.UserRole) func(r *ghttp.Request) {
	roleCode := service.User().GetUserRoleCodeByUserRole(role)
	return func(r *ghttp.Request) {
		bizCtx := service.BizCtx().Get(r.GetCtx())
		if bizCtx == nil || bizCtx.User == nil {
			g.Log("GuardV1").Debugf(r.GetCtx(), "GuardV1: 用户未登录")
			r.Response.Status = http.StatusUnauthorized
			return
		}
		userRoleCode := service.User().GetUserRoleCodeByUserRole(bizCtx.User.Role)
		if userRoleCode < roleCode {
			g.Log("GuardV1").Debugf(r.GetCtx(), "GuardV1: 用户权限不足，需要 %s", role)
			r.Response.Status = http.StatusForbidden
			return
		}
		r.Middleware.Next()
	}
}
