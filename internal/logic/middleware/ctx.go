package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

// Ctx 用于注入业务上下文，实现单次请求中的业务数据共享
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{}
	service.BizCtx().Init(r, customCtx)
	r.Middleware.Next()
}
