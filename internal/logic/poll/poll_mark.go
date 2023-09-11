package poll

import (
	"context"
	"strconv"
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

// GetPollMarksByPollID 获取指定投票的标签列表（不带用户信息）
func (s *sPoll) GetPollMarksByPollID(ctx context.Context, pid uint) ([]int, error) {
	marks, err := dao.PollMarkRelation.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10, // 10 分钟
		Name:     "poll_marks:pid:" + strconv.Itoa(int(pid)),
		Force:    false,
	}).
		Where(dao.PollMarkRelation.Columns().PollId, pid).
		Fields(dao.PollMarkRelation.Columns().MarkId).
		Distinct().
		Array()
	if err != nil {
		return nil, err
	}
	if len(marks) == 0 {
		return make([]int, 0), nil
	}
	return gconv.Ints(marks), nil
}

func (s *sPoll) GetPollMarksByPollIDAndUserID(ctx context.Context, pid, userID int) ([]int, error) {
	marks, err := dao.PollMarkRelation.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10, // 10 分钟
		Name:     "poll_marks:pid:" + strconv.Itoa(pid) + ":uid:" + strconv.Itoa(userID),
		Force:    false,
	}).
		Where(dao.PollMarkRelation.Columns().PollId, pid).
		Where(dao.PollMarkRelation.Columns().UserId, userID).
		Fields(dao.PollMarkRelation.Columns().MarkId).
		Distinct().
		Array()
	if err != nil {
		return nil, err
	}

	if len(marks) == 0 {
		return make([]int, 0), nil
	}
	return gconv.Ints(marks), nil
}
