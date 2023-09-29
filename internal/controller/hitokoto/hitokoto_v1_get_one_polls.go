package hitokoto

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/samber/lo"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"golang.org/x/sync/errgroup"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"

	v1 "github.com/hitokoto-osc/reviewer/api/hitokoto/v1"
)

func (c *ControllerV1) GetOnePolls(ctx context.Context, req *v1.GetOnePollsReq) (res *v1.GetOnePollsRes, err error) {
	results, count, err := dao.Poll.Ctx(ctx).
		Where(dao.Poll.Columns().SentenceUuid, req.UUID).
		Fields(dao.Poll.Columns().Id).
		Page(req.Page, req.PageSize).
		AllAndCount(true)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取投票列表失败")
	}
	g.Log().Debug(ctx, results)
	if len(results) == 0 {
		return &v1.GetOnePollsRes{
			Total:      count,
			Page:       req.Page,
			PageSize:   req.PageSize,
			Collection: make([]model.PollListElement, 0),
		}, gerror.NewCode(gcode.CodeNotFound)
	}
	pollIDs := lo.Map(results, func(item gdb.Record, index int) int {
		return gconv.Int(item["id"])
	})
	eg, egCtx := errgroup.WithContext(ctx)
	polls := make([]model.PollListElement, len(pollIDs))
	for i, pollID := range pollIDs {
		i, pollID := i, pollID
		eg.Go(func() error {
			poll, err := service.Poll().GetPollDetailByID(egCtx, uint(pollID), req.WithPolledData)
			if err != nil {
				return err
			}
			polls[i] = *poll
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	res = &v1.GetOnePollsRes{
		Total:      count,
		Page:       req.Page,
		PageSize:   req.PageSize,
		Collection: polls,
	}
	return res, nil
}
