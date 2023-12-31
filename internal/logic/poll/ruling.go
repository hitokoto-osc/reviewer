package poll

import (
	"context"
	"math"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

func (s *sPoll) GetRulingThreshold(isExpandedPoll bool, totalTickets int) int {
	var threshold int
	if isExpandedPoll {
		if totalTickets < consts.PollRulingNeedForCommonUserPollThreshold {
			threshold = 100000 // 设置一个不可能达到的值
		} else {
			threshold = int(math.Floor(float64(totalTickets) * consts.PollRulingNeedForCommonUserPollRate))
		}
	} else {
		if totalTickets < consts.PollRulingNormalThreshold {
			threshold = consts.PollRulingInitThreshold
		} else {
			threshold = int(math.Floor(float64(totalTickets) * consts.PollRulingNormalRate))
		}
	}
	return threshold
}

// DoRuling 处理投票
//
//nolint:gocyclo
func (s *sPoll) DoRuling(
	ctx context.Context,
	poll *entity.Poll,
	target consts.PollStatus,
) error {
	if target != consts.PollStatusRejected && target != consts.PollStatusApproved && target != consts.PollStatusNeedModify { //nolint:lll
		return gerror.New("无效的投票状态")
	}

	reviewerUID := consts.PollRulingUserID
	user := service.BizCtx().GetUser(ctx)
	if user != nil {
		reviewerUID = int(user.Id) // 如果是管理员投票，则使用管理员的 UID
	}

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
		_, e := dao.Poll.Ctx(ctx).TX(tx).Where(dao.Poll.Columns().Id, poll.Id).Update(do.Poll{
			Status: int(consts.PollStatusProcessing),
		})
		if e != nil {
			return gerror.Wrapf(e, "锁定投票 %d 失败", poll.Id)
		}
		var pending *entity.Pending
		e = dao.Pending.Ctx(ctx).TX(tx).Where(dao.Pending.Columns().Uuid, poll.SentenceUuid).Scan(&pending)
		if e != nil {
			return gerror.Wrapf(e, "获取句子 %s 的 pending 信息失败", poll.SentenceUuid)
		} else if pending == nil {
			return gerror.Newf("获取句子 %s 的 pending 信息失败", poll.SentenceUuid)
		}

		if target == consts.PollStatusApproved || target == consts.PollStatusRejected {
			// 移动句子
			_, e = dao.Pending.Ctx(ctx).TX(tx).Where(dao.Pending.Columns().Uuid, poll.SentenceUuid).Delete()
			if e != nil {
				return gerror.Wrapf(e, "从 pending 删除句子 %s 失败", poll.SentenceUuid)
			}
			if target == consts.PollStatusApproved {
				_, e = dao.Sentence.Ctx(ctx).TX(tx).Unscoped().Insert(do.Sentence{
					Uuid:       pending.Uuid,
					Hitokoto:   pending.Hitokoto,
					Type:       pending.Type,
					From:       pending.From,
					FromWho:    pending.FromWho,
					Creator:    pending.Creator,
					CreatorUid: pending.CreatorUid,
					Reviewer:   reviewerUID,
					CommitFrom: pending.CommitFrom,
					Owner:      pending.Owner,
					CreatedAt:  pending.CreatedAt,
				})
			} else {
				_, e = dao.Refuse.Ctx(ctx).TX(tx).Unscoped().Insert(do.Refuse{
					Uuid:       pending.Uuid,
					Hitokoto:   pending.Hitokoto,
					Type:       pending.Type,
					From:       pending.From,
					FromWho:    pending.FromWho,
					Creator:    pending.Creator,
					CreatorUid: pending.CreatorUid,
					Owner:      pending.Owner,
					Reviewer:   reviewerUID,
					CommitFrom: pending.CommitFrom,
					CreatedAt:  pending.CreatedAt,
				})
			}
			if e != nil {
				return gerror.Wrapf(e, "移动句子 %s 失败", poll.SentenceUuid)
			}
		} else { // 亟待修改
			_, e = dao.Pending.Ctx(ctx).TX(tx).Where(dao.Pending.Columns().Uuid, poll.SentenceUuid).Unscoped().Update(do.Pending{
				PollStatus: int(consts.PollStatusNeedModify),
				Reviewer:   reviewerUID,
			})
			if e != nil {
				return gerror.Wrapf(e, "更新句子 %s 的 pending 信息失败", poll.SentenceUuid)
			}
		}
		// 更新投票状态
		poll.UpdatedAt = gtime.Now()
		poll.Status = int(target)
		_, e = dao.Poll.Ctx(ctx).TX(tx).Where(dao.Poll.Columns().Id, poll.Id).Update(poll)
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

		pollElement := &model.PollElement{
			ID:           uint(poll.Id),
			SentenceUUID: poll.SentenceUuid,
			Sentence: &model.HitokotoV1WithPoll{
				HitokotoV1: model.HitokotoV1{
					ID:         uint(pending.Id),
					UUID:       pending.Uuid,
					Hitokoto:   pending.Hitokoto,
					Type:       consts.HitokotoType(pending.Type),
					From:       pending.From,
					FromWho:    pending.FromWho,
					Creator:    pending.Creator,
					CreatorUID: uint(pending.CreatorUid),
					Reviewer:   uint(reviewerUID),
					Status:     consts.HitokotoStatusPending, // FIXME: 应该和 target 一致
					CreatedAt:  pending.CreatedAt,
				},
				PollStatus: consts.PollStatus(poll.Status),
			},
			Status:             consts.PollStatus(poll.Status),
			Approve:            poll.Accept,
			Reject:             poll.Reject,
			NeedModify:         poll.NeedEdited,
			NeedCommonUserPoll: poll.NeedUserPoll,
			CreatedAt:          (*time.Time)(poll.CreatedAt),
			UpdatedAt:          (*time.Time)(poll.UpdatedAt),
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
			service.Cache().ClearPollUserCache(ctx, uint(pollLog.UserId))
			service.Cache().ClearCacheAfterPollUpdated(ctx, uint(pollLog.UserId), uint(poll.Id), poll.SentenceUuid)
		}

		// DoReviewedNotification
		if err = doNotificationAfterPollRuling(ctx, pollElement, pollLogs); err != nil {
			return gerror.Wrap(err, "发送投票结果通知失败")
		}
		service.Cache().ClearPollListCache(ctx)
		// 提交句子到搜索引擎
		if pollElement.Status != consts.PollStatusNeedModify {
			e = service.Search().AddSentenceToSearch(ctx, pending)
			if e != nil {
				return gerror.Wrap(e, "提交句子到搜索引擎失败")
			}
		}
		return nil
	})
}

func doNotificationAfterPollRuling(
	ctx context.Context,
	pollElement *model.PollElement,
	pollLogs []entity.PollLog,
) error {
	if err := service.Notification().PollFinishedNotification(
		ctx,
		pollElement,
		pollLogs,
	); err != nil {
		return gerror.Wrap(err, "发送投票结果通知失败")
	}
	if pollElement.Status != consts.PollStatusNeedModify {
		if err := doReviewedNotification(ctx, pollElement); err != nil {
			return gerror.Wrap(err, "发送审核结果通知失败")
		}
	}
	return nil
}

func doReviewedNotification(ctx context.Context, pollElement *model.PollElement) error {
	return service.Notification().SentenceReviewedNotification(
		ctx,
		pollElement,
		consts.PollRulingUserID,
		consts.PollRulingUsername,
	)
}
