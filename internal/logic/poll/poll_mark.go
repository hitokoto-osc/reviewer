package poll

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

func (s *sPoll) GetPollMarkLabels(ctx context.Context) ([]entity.PollMark, error) {
	var marks []entity.PollMark
	err := dao.PollMark.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 2, // 2 小时
		Name:     "poll_mark_labels",
		Force:    false,
	}).Scan(&marks)
	return marks, err
}

// GetPollMarksBySentenceUUID 获取指定投票的标签列表（不带用户信息）
func (s *sPoll) GetPollMarksBySentenceUUID(ctx context.Context, uuid string) ([]int, error) {
	marks, err := dao.PollMarkRelation.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10, // 10 分钟
		Name:     "poll_marks:uuid:" + uuid,
		Force:    false,
	}).
		Where(dao.PollMarkRelation.Columns().SentenceUuid, uuid).
		Fields(dao.PollMarkRelation.Columns().MarkId).
		Distinct().
		Array()
	if err != nil {
		return nil, err
	}
	return gconv.Ints(marks), nil
}
