package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
)

func (s *sUser) IncreaseUserPollScore(ctx context.Context, in *model.UserPollScoreInput) error { //nolint:dupl
	if in == nil {
		return gerror.New("参数错误")
	}
	if in.UserID == 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return gerror.New("无法获取用户信息")
		}
		in.UserID = user.Id
	}
	doUpdateScore := func(ctx context.Context, tx gdb.TX) error {
		// 更新用户积分
		_, err := dao.PollUsers.Ctx(ctx).TX(tx).
			Where(dao.PollUsers.Columns().UserId, in.UserID).
			Increment(dao.PollUsers.Columns().Score, in.Score)
		if err != nil {
			return err
		}
		// 创建积分变动记录
		record := do.PollScorePipeline{
			PollId:       in.PollID,
			UserId:       in.UserID,
			SentenceUuid: in.SentenceUUID,
			Type:         "increment",
			Score:        in.Score,
		}
		if in.Reason != "" {
			record.Reason = in.Reason
		}
		_, err = dao.PollScorePipeline.Ctx(ctx).TX(tx).Insert(record)
		return err
	}

	if in.Tx != nil {
		return doUpdateScore(ctx, in.Tx)
	}
	return g.DB().Transaction(ctx, doUpdateScore)
}

func (s *sUser) DecreaseUserPollScore(ctx context.Context, in *model.UserPollScoreInput) error { //nolint:dupl
	if in == nil {
		return gerror.New("参数错误")
	}
	if in.UserID == 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return gerror.New("无法获取用户信息")
		}
		in.UserID = user.Id
	}
	doUpdateScore := func(ctx context.Context, tx gdb.TX) error {
		// 更新用户积分
		_, err := dao.PollUsers.Ctx(ctx).TX(tx).
			Where(dao.PollUsers.Columns().UserId, in.UserID).
			Decrement(dao.PollUsers.Columns().Score, in.Score)
		if err != nil {
			return err
		}
		// 创建积分变动记录
		record := do.PollScorePipeline{
			PollId:       in.PollID,
			UserId:       in.UserID,
			SentenceUuid: in.SentenceUUID,
			Type:         "decrement",
			Score:        in.Score,
		}
		if in.Reason != "" {
			record.Reason = in.Reason
		}
		_, err = dao.PollScorePipeline.Ctx(ctx).TX(tx).Insert(record)
		return err
	}

	if in.Tx != nil {
		return doUpdateScore(ctx, in.Tx)
	}
	return g.DB().Transaction(ctx, doUpdateScore)
}
