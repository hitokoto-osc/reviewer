// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hitokoto

import (
	"context"

	"github.com/hitokoto-osc/reviewer/api/hitokoto/adminV1"
	"github.com/hitokoto-osc/reviewer/api/hitokoto/v1"
)

type IHitokotoAdminV1 interface {
	DeleteList(ctx context.Context, req *adminV1.DeleteListReq) (res *adminV1.DeleteListRes, err error)
	GetList(ctx context.Context, req *adminV1.GetListReq) (res *adminV1.GetListRes, err error)
	MoveSentence(ctx context.Context, req *adminV1.MoveSentenceReq) (res *adminV1.MoveSentenceRes, err error)
	MoveList(ctx context.Context, req *adminV1.MoveListReq) (res *adminV1.MoveListRes, err error)
	UpdateOne(ctx context.Context, req *adminV1.UpdateOneReq) (res *adminV1.UpdateOneRes, err error)
}

type IHitokotoV1 interface {
	GetOnePolls(ctx context.Context, req *v1.GetOnePollsReq) (res *v1.GetOnePollsRes, err error)
	GetHitokoto(ctx context.Context, req *v1.GetHitokotoReq) (res *v1.GetHitokotoRes, err error)
}
