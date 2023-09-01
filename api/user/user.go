// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/api/user/v1"
)

type IUserV1 interface {
	GetApplyToken(ctx context.Context, req *v1.GetApplyTokenReq) (res *v1.GetApplyTokenRes, err error)
	ApplyReviewer(ctx context.Context, req *v1.ApplyReviewerReq) (res *v1.ApplyReviewerRes, err error)
	GetUserPollLog(ctx context.Context, req *v1.GetUserPollLogReq) (res *v1.GetUserPollLogRes, err error)
	GetUserPollResult(ctx context.Context, req *v1.GetUserPollResultReq) (res *v1.GetUserPollResultRes, err error)
	GetUserPollUnreviewed(ctx context.Context, req *v1.GetUserPollUnreviewedReq) (res *v1.GetUserPollUnreviewedRes, err error)
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
	GetUserScoreRecords(ctx context.Context, req *v1.GetUserScoreRecordsReq) (res *v1.GetUserScoreRecordsRes, err error)
}
