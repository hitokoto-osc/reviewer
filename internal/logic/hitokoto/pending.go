package hitokoto

import (
	"context"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

func (s *sHitokoto) GetPendingByUUID(ctx context.Context, uuid string) (hitokoto *entity.Pending, err error) {
	err = dao.Pending.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 10,
		Name:     "pending:uuid:" + uuid,
		Force:    false,
	}).Where(do.Pending{Uuid: uuid}).Scan(&hitokoto)
	return
}

func (s *sHitokoto) TopPendingPollNotOpen(ctx context.Context) (hitokoto *entity.Pending, err error) {
	err = dao.Pending.Ctx(ctx).
		Where(dao.Pending.Columns().PollStatus, consts.PollStatusNotOpen).
		OrderAsc(dao.Pending.Columns().CreatedAt).
		Scan(&hitokoto)
	return
}

func (s *sHitokoto) CountPendingPollNotOpen(ctx context.Context) (count int, err error) {
	count, err = dao.Pending.Ctx(ctx).
		Where(dao.Pending.Columns().PollStatus, consts.PollStatusNotOpen).
		Count()
	return
}
