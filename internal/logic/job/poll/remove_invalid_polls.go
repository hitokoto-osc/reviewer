package poll

import (
	"context"
	"math"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

// RemoveInvalidPolls 清理无效投票
// 基本目的是清理后台干预了的投票
func RemoveInvalidPolls(ctx context.Context) error {
	g.Log().Debug(ctx, "开始清理无效投票...")
	defer g.Log().Debug(ctx, "清理无效投票完成！")
	params := &model.GetPollListInput{
		StatusStart:     0,
		StatusEnd:       2, // 包含 2 是因为：处理系统错误中断导致的状态异常
		WithPollRecords: false,
		WithCache:       false,
		WithMarks:       false,
		PageSize:        100,
		Page:            1,
	}
	totalPage := 0
	for {
		out, err := service.Poll().GetPollList(ctx, params)
		if err != nil {
			return gerror.Wrapf(err, "在处理第 %d 页（每页 %d 条记录）时出现错误：获取记录失败！", params.Page, params.PageSize)
		}
		if totalPage == 0 {
			totalPage = int(math.Ceil(float64(out.Total) / float64(params.PageSize)))
		}
		g.Log().Debugf(ctx, "当前正在处理第 %d 页，每页 %d 记录，共 %d 页（%d 条记录）", params.Page, params.PageSize, totalPage, out.Total)
		for i := range out.Collection {
			poll := &out.Collection[i]
			g.Log().Debugf(ctx, "正在处理投票 %d...", poll.ID)
			if poll.Sentence.Status == consts.HitokotoStatusApproved || poll.Sentence.Status == consts.HitokotoStatusRejected {
				g.Log().Debugf(ctx, "投票 %d: 句子 %s 状态为 %s，属于无效投票（已通过或已驳回），开始清除……", poll.ID, poll.Sentence.UUID, poll.Sentence.Status)
				var status int
				if poll.Sentence.Status == consts.HitokotoStatusApproved {
					status = int(consts.PollStatusApproved)
				} else {
					status = int(consts.PollStatusRejected)
				}
				affectedRows, e := dao.Poll.Ctx(ctx).Where(dao.Poll.Columns().Id, poll.ID).UpdateAndGetAffected(dao.Poll.Columns().Status, status)
				if e != nil {
					return gerror.Wrapf(err, "投票 %d: 更新状态失败！", poll.ID)
				}
				if affectedRows == 0 {
					g.Log().Debugf(ctx, "投票 %d: 更新状态失败，可能是因为已经被其他人处理了。", poll.ID)
					continue
				}
				go service.Cache().ClearCacheAfterPollUpdated(ctx, 0, poll.ID, poll.SentenceUUID) // 清除缓存
			} else {
				g.Log().Debugf(ctx, "投票 %d: 句子 %s 状态为 %s，属于有效投票，跳过……", poll.ID, poll.Sentence.UUID, poll.Sentence.Status)
			}
		}
		params.Page++
		if totalPage < params.Page {
			break
		}
	}
	return nil
}
