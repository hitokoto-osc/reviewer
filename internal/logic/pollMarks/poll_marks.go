package pollMarks

import (
	"context"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

func init() {
	service.RegisterPollMarks(New())
}

type sPollMarks struct {
}

func New() service.IPollMarks {
	return &sPollMarks{}
}

func (s *sPollMarks) List(ctx context.Context) ([]entity.PollMark, error) {
	var marks []entity.PollMark
	err := dao.PollMark.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour * 2, // 2 小时
		Name:     "poll_mark_labels",
		Force:    false,
	}).Scan(&marks)
	return marks, err
}

func (s *sPollMarks) Update(ctx context.Context, mark *entity.PollMark) error {
	if mark == nil || mark.Id == 0 {
		return gerror.New("参数错误")
	}
	defer g.DB().GetCache().Remove(ctx, "SelectCache:poll_mark_labels") //nolint:errcheck
	affected, err := dao.PollMark.Ctx(ctx).Where(dao.PollMark.Columns().Id, mark.Id).
		Data(mark).
		UpdateAndGetAffected()
	if err != nil {
		return gerror.Wrap(err, "修改失败")
	} else if affected == 0 {
		return gerror.New("修改失败：影响行数为 0")
	}
	return nil
}

func (s *sPollMarks) Add(ctx context.Context, mark *do.PollMark) error {
	defer g.DB().GetCache().Remove(ctx, "SelectCache:poll_mark_labels") //nolint:errcheck
	_, err := dao.PollMark.Ctx(ctx).Insert(mark)
	return err
}
