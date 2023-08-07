package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) ApplyReviewer(ctx context.Context, req *v1.ApplyReviewerReq) (res *v1.ApplyReviewerRes, err error) {
	user := service.BizCtx().GetUser(ctx) // User must be not nil, so no need to check.
	tokenInCache, err := gcache.Get(ctx, fmt.Sprintf("user:%d:apply_token", user.Id))
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get token from cache")
	}
	if tokenInCache.IsNil() {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "token not found")
	}
	if tokenInCache.String() != req.VerificationToken {
		return nil, gerror.NewCode(gcode.CodeOperationFailed, "token not match")
	}
	if err = service.User().SetUserRoleReviewer(ctx, user.Id); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to set user role")
	}
	// clear cache
	_, _ = gcache.Remove(ctx, fmt.Sprintf("user:%d:apply_token", user.Id))
	return &v1.ApplyReviewerRes{}, nil
}
