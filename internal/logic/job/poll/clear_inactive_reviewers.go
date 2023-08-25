package poll

import (
	"context"
	"fmt"
	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/dao"
)

func ClearInactiveReviewer(ctx context.Context) error {
	// 获取所有需要清理的用户（30 天未活跃）
	g.Log().Debugf(ctx, "开始清理 %d 天未活跃的审核员……", consts.ReviewerInactiveDays)
	defer g.Log().Debugf(ctx, "清理 %d 天未活跃的审核员完成！", consts.ReviewerInactiveDays)
	records, err := g.DB().Ctx(ctx).Raw(fmt.Sprintf(
		"SELECT `%s` FROM `%s` WHERE `%s` = 1 AND id in (SELECT `%s` FROM `%s` WHERE `%s` < DATE_SUB(now(), INTERVAL %d DAY))",
		dao.Users.Columns().Id,
		dao.Users.Table(),
		dao.Users.Columns().IsReviewer,
		dao.PollUsers.Columns().UserId,
		dao.PollUsers.Table(),
		dao.PollUsers.Columns().UpdatedAt,
		consts.ReviewerInactiveDays,
	)).Array()
	if err != nil {
		return gerror.Wrap(err, "获取需要清理的用户失败")
	}
	g.Log().Debugf(ctx, "共有 %d 个用户需要清理", len(records))
	if len(records) == 0 {
		return nil
	}
	ids := gconv.Uints(records)
	affectedRows, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, ids).UpdateAndGetAffected(dao.Users.Columns().IsReviewer, 0)
	if err != nil {
		return gerror.Wrap(err, "清理失败")
	}
	g.Log().Debugf(ctx, "清理完成，共清理 %d 个用户", affectedRows)
	return nil
}
