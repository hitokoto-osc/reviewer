// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hitokoto

import (
	"context"

	"github.com/hitokoto-osc/reviewer/api/hitokoto/v1"
)

type IHitokotoV1 interface {
	GetHitokoto(ctx context.Context, req *v1.GetHitokotoReq) (res *v1.GetHitokotoRes, err error)
}
