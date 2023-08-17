package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gcfg"

	"github.com/hitokoto-osc/reviewer/utility"

	"github.com/hitokoto-osc/reviewer/internal/controller/hitokoto"

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
		Name:        "reviewer",
		Usage:       "reviewer [OPTIONS]",
		Brief:       "@hitokoto-osc/reviewer 服务端",
		Description: "reviewer 服务端，用于提供 API 服务。",
		Examples: `
			Dev:
				./reviewer

			Test:
				./reviewer -c config.test.yaml
				or
				GF_GCFG_FILE=config.test.yaml GF_GERROR_BRIEF=true ./reviewer

			Prod:
				./reviewer -c config.prod.yaml
				or
				GF_GCFG_FILE=config.prod.yaml GF_GERROR_BRIEF=true ./reviewer`,
		Additional: "更多信息请访问 https://github.com/hitokoto-osc/reviewer",
		Arguments: []gcmd.Argument{
			{
				Name:   "version",
				Short:  "v",
				Brief:  "print version info",
				IsArg:  false,
				Orphan: true,
			},
			{
				Name:   "config",
				Short:  "c",
				Brief:  "config file (default config.yaml)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			if parser.GetOpt("version") != nil {
				utility.PrintVersionInfo()
				return nil
			}
			config := parser.GetOpt("config").String()
			if config != "" {
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(config)
			}
			// GDB SQL 缓存使用 Redis
			// adapter := gcache.NewAdapterRedis(g.Redis())
			// g.DB().GetCache().SetAdapter(adapter)

			// 注册计划任务
			err = service.Job().Register(ctx)
			if err != nil {
				return gerror.Wrap(err, "注册计划任务失败")
			}

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
					group.Bind(user.NewV1(), hitokoto.NewV1())
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
