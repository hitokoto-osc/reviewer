// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package poll

import (
	"context"

	"github.com/hitokoto-osc/reviewer/api/poll/adminV1"
	"github.com/hitokoto-osc/reviewer/api/poll/v1"
)

type IPollAdminV1 interface {
	AddOneMark(ctx context.Context, req *adminV1.AddOneMarkReq) (res *adminV1.AddOneMarkRes, err error)
	UpdateOneMark(ctx context.Context, req *adminV1.UpdateOneMarkReq) (res *adminV1.UpdateOneMarkRes, err error)
}

type IPollV1 interface {
	CancelPoll(ctx context.Context, req *v1.CancelPollReq) (res *v1.CancelPollRes, err error)
	GetPollDetail(ctx context.Context, req *v1.GetPollDetailReq) (res *v1.GetPollDetailRes, err error)
	GetPolls(ctx context.Context, req *v1.GetPollsReq) (res *v1.GetPollsRes, err error)
	GetPollMarks(ctx context.Context, req *v1.GetPollMarksReq) (res *v1.GetPollMarksRes, err error)
	NewPoll(ctx context.Context, req *v1.NewPollReq) (res *v1.NewPollRes, err error)
	Poll(ctx context.Context, req *v1.PollReq) (res *v1.PollRes, err error)
}
