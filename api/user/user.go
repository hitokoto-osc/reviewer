// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/api/user/v1"
)

type IUserV1 interface {
	GetUserPollLog(ctx context.Context, req *v1.GetUserPollLogReq) (res *v1.GetUserPollLogRes, err error)
	GetUserPollResult(ctx context.Context, req *v1.GetUserPollResultReq) (res *v1.GetUserPollResultRes, err error)
	GetUser(ctx context.Context, req *v1.GetUserReq) (res *v1.GetUserRes, err error)
}
