package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"golang.org/x/sync/errgroup"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserPollResult(ctx context.Context, req *v1.GetUserPollResultReq) (res *v1.GetUserPollResultRes, err error) {
	var (
		count int
		out   *model.GetUserPollLogsWithPollResultOutput
	)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		count, err = service.User().GetUserPollLogsCount(egCtx, 0)
		return err
	})
	eg.Go(func() error {
		out, err = service.User().GetUserPollLogsWithPollResult(egCtx, model.GetUserPollLogsWithPollResultInput{
			Order:     dao.PollLog.Columns().CreatedAt + " " + req.Order,
			Page:      req.Page,
			PageSize:  req.PageSize,
			WithCache: true,
		})
		return err
	})
	if err = eg.Wait(); err != nil {
		return nil, gerror.Wrap(err, "get user poll result failed")
	}
	res = &v1.GetUserPollResultRes{
		Total:                               uint(count),
		GetUserPollLogsWithPollResultOutput: *out,
	}

	return res, nil
}
