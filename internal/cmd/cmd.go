package cmd

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/controller/admin"

	"github.com/hitokoto-osc/reviewer/internal/controller/poll"
	"github.com/hitokoto-osc/reviewer/internal/controller/user"

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
			s.AddSearchPath("resource/public")
			s.Use(service.Middleware().HandlerResponse)
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().AuthorizationV1)
				group.Bind(index.NewV1(), user.NewV1(), poll.NewV1())
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().AuthorizationAdminV1)
					group.Bind(admin.NewV1())
				})
			})
			s.Run()
			return nil
		},
	}
)
