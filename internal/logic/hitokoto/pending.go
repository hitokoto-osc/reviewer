package hitokoto

import (
	"context"
	"time"

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
