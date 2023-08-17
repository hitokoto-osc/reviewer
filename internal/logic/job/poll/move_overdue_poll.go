package poll

import (
	"context"
	"math"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/hitokoto-osc/reviewer/internal/model/do"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/dao"
)

// MoveOverduePolls 移动过期投票
// 将过期投票直接修改为亟待审核
func MoveOverduePolls(ctx context.Context) error {
	g.Log().Debug(ctx, "开始移动过期投票...")
	var (
		polls     []entity.Poll
		page      = 1
		pageSize  = 100
		total     int
		totalPage int
	)
	query := dao.Poll.Ctx(ctx).
		Where(dao.Poll.Columns().Status, consts.PollStatusOpen).
		WhereLT(dao.Poll.Columns().CreatedAt, gdb.Raw("INTERVAL 15 DAY")) // 15 天
	total, err := query.Clone().Fields("1").Count()
	if err != nil {
		return gerror.Wrap(err, "获取过期投票数量失败")
	}
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
		var (
			poll    *entity.Poll
			userIDs []uint
		)
		for i := 0; i < len(polls); i++ {
			poll = &polls[i]
			pollLogs, e := service.Poll().GetPollLogsByPollID(ctx, poll.Id)
			if e != nil {
				return gerror.Wrapf(e, "获取投票 %d 的投票记录失败", poll.Id)
			}
			if len(pollLogs) == 0 {
				g.Log().Debugf(ctx, "投票 %d 没有投票记录，跳过", poll.Id)
				continue
			}
			userIDs = make([]uint, len(pollLogs))
			for j := 0; j < len(pollLogs); j++ {
				userIDs[j] = uint(pollLogs[j].UserId)
			}
		}
		// 事务更新结果
		e := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
			// 更新投票状态
			_, e := dao.Poll.Ctx(ctx).TX(tx).Where(dao.Poll.Columns().Id, poll.Id).Update(dao.Poll.Columns().Status, consts.PollStatusNeedModify)
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
			for userID := range userIDs {
				e = service.User().IncreaseUserPollScore(ctx, &model.UserPollScoreInput{
					UserID:       uint(userID),
					PollID:       uint(poll.Id),
					Score:        consts.PollParticipantScore,
					SentenceUUID: poll.SentenceUuid,
					Reason:       "感谢参与，投票超时自动处理",
					Tx:           tx,
				})
				if e != nil {
					return gerror.Wrap(e, "赋予用户积分失败")
				}
			}
			return nil
		})
		if e != nil {
			return gerror.Wrap(e, "更新投票状态失败")
		}
		page++
		if page > totalPage {
			break
		}
	}
	g.Log().Debug(ctx, "移动过期投票完成")
	return nil
}
