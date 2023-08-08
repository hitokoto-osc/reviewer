package poll

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"time"
)

func (s *sPoll) GetPollLogsBySentenceUUID(ctx context.Context, uuid string) ([]entity.PollLog, error) {
	var logs []entity.PollLog
	err := dao.PollLog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 2,
		Name:     "poll_log:uuid:" + uuid,
		Force:    false,
	}).Where(dao.PollLog.Columns().SentenceUuid, uuid).Scan(&logs)
	return logs, err
}
