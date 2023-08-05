// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package index

import (
	"context"

	"github.com/hitokoto-osc/reviewer/api/index/common"
	"github.com/hitokoto-osc/reviewer/api/index/v1"
)

type IIndexCommon interface {
	Index(ctx context.Context, req *common.IndexReq) (res *common.IndexRes, err error)
}

type IIndexV1 interface {
	GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error)
}
