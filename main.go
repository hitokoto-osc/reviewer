package main

import (
	_ "reviewer/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"reviewer/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
