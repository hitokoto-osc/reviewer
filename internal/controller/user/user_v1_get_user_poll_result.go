package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserPollResult(ctx context.Context, req *v1.GetUserPollResultReq) (res *v1.GetUserPollResultRes, err error) { // nolint: lll
	out, err := service.User().GetUserPollLogsWithPollResult(ctx, model.GetUserPollLogsWithPollResultInput{
		Order:     dao.PollLog.Columns().CreatedAt + " " + req.Order,
		Page:      req.Page,
		PageSize:  req.PageSize,
		WithCache: true,
	})

	if err != nil {
		return nil, gerror.Wrap(err, "get user poll result failed")
	}
	res = &v1.GetUserPollResultRes{
		Total:                               uint(out.Total),
		GetUserPollLogsWithPollResultOutput: *out,
	}

	return res, nil
}
