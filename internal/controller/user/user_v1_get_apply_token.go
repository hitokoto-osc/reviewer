package user

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/gogf/gf/v2/os/gtime"
	vtime "github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/google/uuid"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetApplyToken(ctx context.Context, req *v1.GetApplyTokenReq) (res *v1.GetApplyTokenRes, err error) {
	user := service.BizCtx().GetUser(ctx) // User must be not nil, so no need to check.
	if service.User().GetUserRoleCodeByUserRole(user.Role) >= consts.UserRoleCodeReviewer {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "user already reviewer")
	}
	// Anyway, we will return a new token.
	token, err := uuid.NewRandom()
	if err != nil {
		return nil, gerror.Wrapf(err, "failed to generate token")
	}
	// Set token to cache.
	ex := time.Minute * 10
	err = gcache.Set(ctx, fmt.Sprintf("user:%d:apply_token", user.Id), token.String(), ex)
	if err != nil {
		return nil, gerror.Wrapf(err, "failed to set token to cache")
	}
	res = &v1.GetApplyTokenRes{
		VerificationToken: token.String(),
		ValidUntil:        (*vtime.Time)(gtime.Now().Add(ex)),
	}
	return res, nil
}
