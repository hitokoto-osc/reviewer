package index

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/hitokoto-osc/reviewer/api/index/common"
)

func (c *ControllerCommon) Index(ctx context.Context, req *common.IndexReq) (res *common.IndexRes, err error) {
	res = &common.IndexRes{
		Name:      consts.AppName,
		Version:   consts.Version,
		Donate:    "喜欢我们的服务吗？赞助我们： https://hitokoto.cn/donate",
		Website:   "https://hitokoto.cn/",
		Copyright: "MoeTeam © 2016-2023 All Rights Reserved.",
	}
	return
}
