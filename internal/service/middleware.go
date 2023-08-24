// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type (
	IMiddleware interface {
		// AuthorizationV1 用于 v1 接口校验用户是否登录
		// 尝试顺序 Authorization: Bearer Token -> param -> form -> body -> query -> Router
		AuthorizationV1(r *ghttp.Request)
		// Ctx 用于注入业务上下文，实现单次请求中的业务数据共享
		Ctx(r *ghttp.Request)
		// GuardV1 是 v1 用于检查用户权限的中间件
		GuardV1(role consts.UserRole) func(r *ghttp.Request)
		CORS(r *ghttp.Request)
		// HandlerResponse 重写了默认的 JSON 响应格式，提供统一的响应格式
		HandlerResponse(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
