package user

import (
	"context"
	"fmt"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	vtime "github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"golang.org/x/sync/errgroup"

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

func (s *sUser) CountUserScoreRecords(ctx context.Context, userID uint) (int, error) {
	return dao.PollScorePipeline.Ctx(ctx).Where(dao.PollScorePipeline.Columns().UserId, userID).Count()
}

func (s *sUser) GetUserScoreRecords(
	ctx context.Context,
	in *model.GetUserScoreRecordsInput,
) (*model.GetUserScoreRecordsOutput, error) {
	if in == nil {
		return nil, gerror.New("参数错误")
	}
	if in.UserID == 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.New("无法获取用户信息")
		}
		in.UserID = user.Id
	}
	out := &model.GetUserScoreRecordsOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
	}
	query := dao.PollScorePipeline.Ctx(ctx).Where(dao.PollScorePipeline.Columns().UserId, in.UserID).Page(in.Page, in.PageSize)
	if in.Order != "" {
		query = query.Order(in.Order)
	}
	if in.WithCache {
		query = query.Cache(gdb.CacheOption{
			Duration: time.Minute * 5,
			Name:     fmt.Sprintf("user:score:records:uid:%d:page:%d:pageSize:%d:order:%s", in.UserID, in.Page, in.PageSize, in.Order),
			Force:    false,
		})
	}

	var records []entity.PollScorePipeline
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var err error
		out.Total, err = s.CountUserScoreRecords(egCtx, in.UserID)
		return err
	})
	eg.Go(func() error {
		return query.Ctx(egCtx).Scan(&records)
	})
	if err := eg.Wait(); err != nil {
		return nil, gerror.Wrap(err, "获取用户积分记录失败")
	}
	// convert to output
	out.Collection = make([]model.UserScoreRecord, len(records))
	for i, v := range records {
		out.Collection[i] = model.UserScoreRecord{
			ID:           uint(v.Id),
			PollID:       uint(v.PollId),
			UserID:       uint(v.UserId),
			SentenceUUID: v.SentenceUuid,
			Score:        v.Score,
			Type:         consts.UserScoreType(v.Type),
			Reason:       v.Reason,
			UpdatedAt:    (*vtime.Time)(v.UpdatedAt),
			CreatedAt:    (*vtime.Time)(v.CreatedAt),
		}
	}
	return out, nil
}
