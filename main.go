package main

import (
	_ "poll/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"poll/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
