package utility

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

func PrintVersionInfo() {
	info := gbuild.Info()
	if info.Git == "" {
		info.Git = "none"
	}
	fmt.Printf(`%s %s

Git Commit:  %s
Build Time:  %s
Go Version:  %s
GF Version:  %s

MoeTeam © 2023 All Rights Reserved.
`,
		consts.AppName,
		consts.Version,
		info.Git,
		info.Time,
		info.Golang,
		info.GoFrame,
	)
}
