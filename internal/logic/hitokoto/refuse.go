package hitokoto

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

func (s *sHitokoto) GetRefuseByUUID(ctx context.Context, uuid string) (hitokoto *entity.Refuse, err error) {
	err = dao.Refuse.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 10,
		Name:     "refuse:uuid:" + uuid,
		Force:    false,
	}).Where(do.Refuse{Uuid: uuid}).Scan(&hitokoto)
	return
}
