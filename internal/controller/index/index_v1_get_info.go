package index

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/hitokoto-osc/reviewer/api/index/v1"
)

func (c *ControllerV1) GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error) {
	res = &v1.GetInfoRes{
		Name:    consts.AppName,
		Version: consts.APIVersionV1,
		Donate:  "喜欢我们的服务吗？赞助我们： https://hitokoto.cn/donate",
		Website: "https://hitokoto.cn/",
		Feedback: g.MapStrStr{
			"desc":        "向我们发信反馈问题吧！",
			"kuertianshi": "i@loli.online",
			"freejishu":   "i@freejishu.com",
			"a632079":     "a632079@qq.com",
		},
		Copyright: "MoeTeam © 2016-2023 All Rights Reserved.",
	}
	return
}
