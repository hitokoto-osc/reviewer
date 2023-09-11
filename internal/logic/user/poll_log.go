package user

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	vtime "github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

func (s *sUser) GetUserPollLogs(ctx context.Context, in model.GetUserPollLogsInput) (*model.GetUserPollLogsOutput, error) {
	if in.UserID <= 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.New("user not found")
		}
		in.UserID = user.Id
	}
	query := dao.PollLog.Ctx(ctx)
	out := &model.GetUserPollLogsOutput{}
	if in.WithCache {
		query = query.Cache(gdb.CacheOption{
			Duration: time.Second * 3,
			Name:     fmt.Sprintf("poll_logs:uid:%d:page:%d:pageSize:%d:order:%s", in.UserID, in.Page, in.PageSize, in.Order),
			Force:    false,
		})
	}
	query = query.Where(dao.PollLog.Columns().UserId, in.UserID)
	if in.Page > 0 && in.PageSize > 0 {
		query = query.Page(in.Page, in.PageSize)
		out.Page, out.PageSize = in.Page, in.PageSize
	}

	if in.Order != "" {
		query = query.Order(in.Order)
	}
	// fetch data from query
	var pollLogs []entity.PollLog
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var e error
		out.Total, e = s.GetUserPollLogsCount(egCtx, in.UserID)
		return e
	})
	eg.Go(func() error {
		return query.Scan(&pollLogs)
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	userPollLogs := make([]model.UserPollLog, len(pollLogs))
	for i, v := range pollLogs {
		userPollLogs[i] = model.UserPollLog{
			PollID:       uint(v.PollId),
			Point:        v.Point,
			SentenceUUID: v.SentenceUuid,
			Method:       consts.PollMethod(v.Type),
			Comment:      v.Comment,
			CreatedAt:    (*vtime.Time)(v.CreatedAt),
			UpdatedAt:    (*vtime.Time)(v.UpdatedAt),
		}
	}
	out.Collection = userPollLogs
	return out, nil
}

func (s *sUser) GetUserPollLogsWithSentence(
	ctx context.Context,
	in model.GetUserPollLogsInput,
) (*model.GetUserPollLogsWithSentenceOutput, error) {
	if in.UserID <= 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.New("user not found")
		}
		in.UserID = user.Id
	}
	pollLogs, err := s.GetUserPollLogs(ctx, in)
	if err != nil {
		return nil, err
	}
	out := &model.GetUserPollLogsWithSentenceOutput{
		Page:     in.Page,
		PageSize: in.PageSize,
		Total:    pollLogs.Total,
	}
	userPollLogs := make([]model.UserPollLogWithSentenceAndUserMarks, len(pollLogs.Collection))
	// 并发获取句子
	eg, egCtx := errgroup.WithContext(ctx)
	for i, v := range pollLogs.Collection {
		index, value := i, v // 解决并发问题
		eg.Go(func() error {
			var e error
			userPollLogs[index].UserMarks, e = service.Poll().GetPollMarksByPollIDAndUserID(egCtx, int(value.PollID), int(in.UserID))
			return e
		})
		eg.Go(func() error {
			sentence, e := service.Hitokoto().GetHitokotoV1SchemaByUUID(egCtx, value.SentenceUUID)
			if e != nil {
				return e
			}
			userPollLogs[index].UserPollLog = value
			userPollLogs[index].Sentence = sentence
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		return nil, err
	}

	out.Collection = userPollLogs
	return out, nil
}

// GetUserPollLogsCount 获取用户投票记录数量
func (s *sUser) GetUserPollLogsCount(ctx context.Context, userID uint) (int, error) {
	if userID <= 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return 0, gerror.New("user not found")
		}
		userID = user.Id
	}
	return dao.PollLog.Ctx(ctx).Where(dao.PollLog.Columns().UserId, userID).Count()
}

func (s *sUser) GetUserPollLogsWithPollResult(
	ctx context.Context,
	in model.GetUserPollLogsInput,
) (*model.GetUserPollLogsWithPollResultOutput, error) {
	pollLogs, err := s.GetUserPollLogsWithSentence(ctx, in)
	if err != nil {
		return nil, err
	}
	collections := make([]model.UserPollElement, len(pollLogs.Collection))
	// Copy Properties and fetch poll result
	eg, egCtx := errgroup.WithContext(ctx)
	for i, v := range pollLogs.Collection {
		index, value := i, v
		collections[index] = model.UserPollElement{
			UserPollLogWithSentenceAndUserMarks: value,
		}
		// 并发获取投票结果
		eg.Go(func() error {
			poll, e := service.Poll().GetPollByID(egCtx, int(value.PollID))
			if e != nil {
				return e
			}
			collections[index].PollInfo = &model.PollElement{
				SentenceUUID:       poll.SentenceUuid,
				Status:             consts.PollStatus(poll.Status),
				Approve:            poll.Accept,
				Reject:             poll.Reject,
				NeedModify:         poll.NeedEdited,
				NeedCommonUserPoll: poll.NeedUserPoll,
				CreatedAt:          (*vtime.Time)(poll.CreatedAt),
				UpdatedAt:          (*vtime.Time)(poll.UpdatedAt),
			}
			return nil
		})
		// 并发获取标记
		eg.Go(func() error {
			marks, e := service.Poll().GetPollMarksByPollID(egCtx, value.PollID)
			if e != nil {
				return e
			}

			collections[index].PollMarks = marks
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		return nil, err
	}
	out := &model.GetUserPollLogsWithPollResultOutput{
		Collection: collections,
		Page:       in.Page,
		PageSize:   in.PageSize,
		Total:      pollLogs.Total,
	}
	return out, nil
}

func (s *sUser) GetUserPolledDataWithPollID(ctx context.Context, userID, pid uint) (*model.PolledData, error) {
	if userID <= 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.New("user not found")
		}
		userID = user.Id
	}
	var pollLog *entity.PollLog
	err := dao.PollLog.Ctx(ctx).
		Where(dao.PollLog.Columns().PollId, pid).
		Where(dao.PollLog.Columns().UserId, userID).
		WhereNot(dao.PollLog.Columns().Type, consts.PollMethodNeedCommonUserPoll). // 排除需要普通用户投票的记录
		Scan(&pollLog)
	if err != nil {
		return nil, err
	} else if pollLog == nil {
		return nil, nil
	}
	res := &model.PolledData{
		Point:     pollLog.Point,
		Method:    consts.PollMethod(pollLog.Type),
		CreatedAt: (*vtime.Time)(pollLog.CreatedAt),
		UpdatedAt: (*vtime.Time)(pollLog.UpdatedAt),
	}
	return res, nil
}
