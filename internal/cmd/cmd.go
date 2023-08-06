package cmd

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/controller/poll"
	"github.com/hitokoto-osc/reviewer/internal/controller/user"

	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/hitokoto-osc/reviewer/internal/controller/admin"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/hitokoto-osc/reviewer/internal/controller/index"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	_ "github.com/hitokoto-osc/reviewer/internal/boot"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetServerAgent(consts.AppName + " " + consts.Version) // 设置服务名称
			s.AddSearchPath("resource/public")                      // 静态文件
			s.Use(service.Middleware().HandlerResponse)             // 统一返回格式
			s.BindHandler("/", index.NewCommon().Index)             // 首页
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Bind(index.NewV1())
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().Ctx,
						service.Middleware().AuthorizationV1,
					)
					group.Bind(user.NewV1())
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(service.Middleware().GuardV1(consts.UserRoleReviewer))
						group.Bind(poll.NewV1())
					})
					group.Group("/admin", func(group *ghttp.RouterGroup) {
						group.Middleware(service.Middleware().GuardV1(consts.UserRoleAdmin))
						group.Bind(admin.NewV1())
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
