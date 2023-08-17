package poll

import (
	"context"
	"math"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/hitokoto-osc/reviewer/internal/consts"
)

func (s *sPoll) GetRulingThreshold(isExpandedPoll bool, totalTickets int) int {
	var threshold int
	if isExpandedPoll {
		if totalTickets < consts.PollRulingNeedForCommonUserPollThreshold {
			threshold = 100000 // 设置一个不可能达到的值
		} else {
			threshold = int(math.Round(float64(totalTickets) * consts.PollRulingNeedForCommonUserPollRate))
		}
	} else {
		if totalTickets < consts.PollRulingNormalThreshold {
			threshold = consts.PollRulingInitThreshold
		} else {
			threshold = int(math.Round(float64(totalTickets) * consts.PollRulingNormalRate))
		}
	}
	return threshold
}

func (s *sPoll) DoRuling(ctx context.Context, poll *entity.Poll, target consts.PollStatus) error {
	// 判断是否需要修改
	var (
		pollLogs []entity.PollLog
		err      error
	)
	pollLogs, err = service.Poll().GetPollLogsByPollID(ctx, poll.Id)
	if err != nil {
		return gerror.Wrapf(err, "获取投票 %d 的投票记录失败", poll.Id)
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 首先先锁定投票
		_, e := dao.Poll.Ctx(ctx).TX(tx).Where(dao.Poll.Columns().Id, poll.Id).LockUpdate().Update(do.Poll{
			Status: int(consts.PollStatusProcessing),
		})
		if e != nil {
			return gerror.Wrapf(e, "锁定投票 %d 失败", poll.Id)
		}
		var pending *entity.Pending
		e = dao.Pending.Ctx(ctx).TX(tx).Where(dao.Pending.Columns().Uuid, poll.SentenceUuid).Scan(&pending)
		if e != nil {
			return gerror.Wrapf(e, "获取句子 %s 的 pending 信息失败", poll.SentenceUuid)
		}
		if target == consts.PollStatusApproved || target == consts.PollStatusRejected {
			// 移动句子
			_, e = dao.Pending.Ctx(ctx).TX(tx).Where(dao.Pending.Columns().Uuid, poll.SentenceUuid).Delete()
			if e != nil {
				return gerror.Wrapf(e, "从 pending 删除句子 %s 失败", poll.SentenceUuid)
			}
			if target == consts.PollStatusApproved {
				_, e = dao.Sentence.Ctx(ctx).TX(tx).Insert(do.Sentence{
					Uuid:       pending.Uuid,
					Hitokoto:   pending.Hitokoto,
					Type:       pending.Type,
					From:       pending.From,
					FromWho:    pending.FromWho,
					Creator:    pending.Creator,
					CreatorUid: pending.CreatorUid,
					Reviewer:   consts.PollRulingUserID,
					CommitFrom: pending.CommitFrom,
					Owner:      pending.Owner,
					CreatedAt:  pending.CreatedAt,
				})
			} else {
				_, e = dao.Refuse.Ctx(ctx).TX(tx).Insert(do.Refuse{
					Uuid:       pending.Uuid,
					Hitokoto:   pending.Hitokoto,
					Type:       pending.Type,
					From:       pending.From,
					FromWho:    pending.FromWho,
					Creator:    pending.Creator,
					CreatorUid: pending.CreatorUid,
					Owner:      pending.Owner,
					Reviewer:   consts.PollRulingUserID,
					CommitFrom: pending.CommitFrom,
					CreatedAt:  pending.CreatedAt,
				})
			}
			if e != nil {
				return gerror.Wrapf(e, "移动句子 %s 失败", poll.SentenceUuid)
			}
		} else { // 亟待修改
			_, e = dao.Pending.Ctx(ctx).TX(tx).Where(dao.Pending.Columns().Uuid, poll.SentenceUuid).Update(do.Pending{
				PollStatus: int(consts.PollStatusNeedModify),
				Reviewer:   consts.PollRulingUserID,
			})
			if e != nil {
				return gerror.Wrapf(e, "更新句子 %s 的 pending 信息失败", poll.SentenceUuid)
			}
		}
		// 更新投票状态
		_, e = dao.Poll.Ctx(ctx).TX(tx).Where(dao.Poll.Columns().Id, poll.Id).Update(do.Poll{
			Status: int(target),
		})
		if e != nil {
			return gerror.Wrapf(e, "更新投票 %d 状态失败", poll.Id)
		}
		// 新增操作日记
		_, e = dao.PollPipeline.Ctx(ctx).TX(tx).Insert(do.PollPipeline{
			PollId:       poll.Id,
			SentenceUuid: poll.SentenceUuid,
			Operate:      int(target),
		})
		if e != nil {
			return gerror.Wrapf(e, "新增操作日记失败")
		}
		// 赋予用户积分
		for _, pollLog := range pollLogs {
			var score int
			if consts.PollMethod(pollLog.Type) == s.translatePollStatusToMethod(target) {
				score = consts.PollWinnerScore
			} else {
				score = consts.PollParticipantScore
			}

			e = service.User().IncreaseUserPollScore(ctx, &model.UserPollScoreInput{
				UserID:       uint(pollLog.UserId),
				PollID:       uint(poll.Id),
				Score:        score,
				SentenceUUID: poll.SentenceUuid,
				Tx:           tx,
			})
			if e != nil {
				return gerror.Wrap(e, "赋予用户积分失败")
			}
		}
		return nil
	})
}
