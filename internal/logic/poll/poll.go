package poll

import (
	"context"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type sPoll struct{}

func init() {
	service.RegisterPoll(New())
}

func New() service.IPoll {
	return &sPoll{}
}

func (s *sPoll) GetPollBySentenceUUID(ctx context.Context, uuid string) (poll *entity.Poll, err error) {
	err = dao.Poll.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10,
		Name:     "poll:uuid:" + uuid,
		Force:    false,
	}).Where(do.Poll{SentenceUuid: uuid}).Scan(&poll)
	return
}
