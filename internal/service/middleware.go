// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// HandlerResponse 重写了默认的 JSON 响应格式，提供统一的响应格式
		HandlerResponse(r *ghttp.Request)
		// AuthorizationV1 用于 v1 接口校验用户是否登录且是否具有审核员权限
		AuthorizationV1(r *ghttp.Request)
		// AuthorizationAdminV1 用于 v1 接口校验用户是否登录且是否具有管理员权限
		AuthorizationAdminV1(r *ghttp.Request)
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
