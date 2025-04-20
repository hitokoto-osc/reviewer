package poll

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"math"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/dao"
)

// MoveOverduePolls 移动过期投票
// 将过期投票直接修改为亟待审核
func MoveOverduePolls(ctx context.Context) error {
	g.Log().Debug(ctx, "开始移动过期投票...")
	defer g.Log().Debug(ctx, "移动过期投票任务执行完成")
	var (
		polls     []entity.Poll
		page      = 1
		pageSize  = 100
		total     int
		totalPage int
	)
	query := dao.Poll.Ctx(ctx).
		Where(dao.Poll.Columns().Status, consts.PollStatusOpen).
		Wheref("%s < DATE_SUB(now(), INTERVAL %d DAY)", dao.Poll.Columns().CreatedAt, consts.PollOverdueDays) // 15 天
	total, err := query.Clone().Count()
	if err != nil {
		return gerror.Wrap(err, "获取过期投票数量失败")
	}
	g.Log().Debugf(ctx, "共有 %d 个过期投票", total)
	if total == 0 {
		return nil
	}
	totalPage = int(math.Ceil(float64(total) / float64(pageSize)))

	for {
		g.Log().Debugf(ctx, "正在处理第 %d 页，每页 %d 条记录，共 %d 页（%d 条记录）", page, pageSize, totalPage, total)
		err = query.Clone().Page(page, pageSize).Scan(&polls)
		if err != nil {
			return gerror.Wrapf(err, "获取第 %d 页记录失败", page)
		}
		// 获取 pollLogs
		for i := 0; i < len(polls); i++ {
			poll := &polls[i]
			pollLogs, e := service.Poll().GetPollLogsByPollID(ctx, poll.Id)
			if e != nil {
				return gerror.Wrapf(e, "获取投票 %d 的投票记录失败", poll.Id)
			}
			if len(pollLogs) == 0 {
				g.Log().Debugf(ctx, "投票 %d 没有投票记录，跳过", poll.Id)
				continue
			}
			userIDs := make([]uint, len(pollLogs))
			for j := 0; j < len(pollLogs); j++ {
				userIDs[j] = uint(pollLogs[j].UserId)
			}
			g.Log().Debugf(ctx,  "处理投票 %d，涉及用户 %v", poll.Id, userIDs)
			// 事务更新结果
			e = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
				// 更新投票状态
				_, e := dao.Poll.Ctx(ctx).TX(tx).Where(dao.Poll.Columns().Id, poll.Id).Update(do.Poll{
					Status: int(consts.PollStatusNeedModify),
				})
				if e != nil {
					return gerror.Wrap(e, "更新投票状态失败")
				}
				// 新增操作日记
				_, e = dao.PollPipeline.Ctx(ctx).TX(tx).Insert(do.PollPipeline{
					PollId:       poll.Id,
					SentenceUuid: poll.SentenceUuid,
					Operate:      consts.PollStatusNeedModify,
					Mark:         "超时自动处理",
				})
				if e != nil {
					return gerror.Wrap(e, "新增操作日记失败")
				}
				// 赋予用户积分
				for _, userID := range userIDs {
					g.Log().Debugf(ctx, "<UNK> %d <UNK>", userID)
					e = service.User().IncreaseUserPollScore(ctx, &model.UserPollScoreInput{
						UserID: uint(userID),
						PollID:       uint(poll.Id),
						Score:        consts.PollParticipantScore,
						SentenceUUID: poll.SentenceUuid,
						Reason:       "感谢参与，投票超时自动处理",
						Tx:           tx,
					})
					if e != nil {
						return gerror.Wrapf(e, "赋予用户 %d 积分失败", userID)
					}
				}
				return nil
			})
			if e != nil {
				return gerror.Wrapf(e, "更新投票 %d 状态失败", poll.Id)
			}
		}

		page++
		if page > totalPage {
			break
		}
	}
	g.Log().Debug(ctx, "移动过期投票完成")
	return nil
}
