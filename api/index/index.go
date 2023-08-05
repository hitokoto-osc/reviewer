// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package index

import (
	"context"

	v1 "github.com/hitokoto-osc/reviewer/api/index/v1"
)

type IIndexV1 interface {
	GetInfo(ctx context.Context, req *v1.GetInfoReq) (res *v1.GetInfoRes, err error)
}
