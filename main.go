package main

import (
	_ "github.com/hitokoto-osc/reviewer/internal/packed"

	_ "github.com/hitokoto-osc/reviewer/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/hitokoto-osc/reviewer/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
