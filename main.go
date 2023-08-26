package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	_ "github.com/hitokoto-osc/reviewer/internal/boot"
	_ "github.com/hitokoto-osc/reviewer/internal/logic"
	_ "github.com/hitokoto-osc/reviewer/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/hitokoto-osc/reviewer/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
