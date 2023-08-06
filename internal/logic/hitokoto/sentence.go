package hitokoto

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

func (s *sHitokoto) GetSentenceByUUID(ctx context.Context, uuid string) (hitokoto *entity.Sentence, err error) {
	err = dao.Sentence.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 10,
		Name:     "sentence:uuid:" + uuid,
		Force:    false,
	}).Where(do.Sentence{Uuid: uuid}).Scan(&hitokoto)
	return
}
