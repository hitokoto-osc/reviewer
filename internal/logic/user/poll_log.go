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
	err := query.Scan(&pollLogs)
	if err != nil {
		return nil, err
	}
	userPollLogs := make([]model.UserPollLog, len(pollLogs))
	for i, v := range pollLogs {
		userPollLogs[i] = model.UserPollLog{
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
	query := dao.PollLog.Ctx(ctx)
	out := &model.GetUserPollLogsWithSentenceOutput{}
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
	err := query.Scan(&pollLogs)
	if err != nil {
		return nil, err
	}
	userPollLogs := make([]model.UserPollLogWithSentence, len(pollLogs))
	// 并发获取句子
	eg, _ := errgroup.WithContext(ctx)
	for i, v := range pollLogs {
		index, value := i, v // 解决并发问题
		eg.Go(func() error {
			sentence, e := service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, value.SentenceUuid)
			if e != nil {
				return e
			}
			userPollLogs[index] = model.UserPollLogWithSentence{
				UserPollLog: model.UserPollLog{
					Point:        value.Point,
					SentenceUUID: value.SentenceUuid,
					Method:       consts.PollMethod(value.Type),
					Comment:      value.Comment,
					CreatedAt:    (*vtime.Time)(value.CreatedAt),
					UpdatedAt:    (*vtime.Time)(value.UpdatedAt),
				},
				Sentence: sentence,
			}
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
			UserPollLogWithSentence: value,
		}
		// 并发获取投票结果
		eg.Go(func() error {
			poll, e := service.Poll().GetPollBySentenceUUID(egCtx, value.SentenceUUID)
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
			marks, e := service.Poll().GetPollMarksBySentenceUUID(egCtx, value.SentenceUUID)
			if e != nil {
				return e
			}

			collections[index].Marks = marks
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
	}
	return out, nil
}

func (s *sUser) GetUserPolledDataWithSentenceUUID(ctx context.Context, userID uint, sentenceUUID string) (*model.PolledData, error) {
	if userID <= 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.New("user not found")
		}
		userID = user.Id
	}
	var pollLog *entity.PollLog
	err := dao.PollLog.Ctx(ctx).
		Where(dao.PollLog.Columns().SentenceUuid, sentenceUUID).
		Where(dao.PollLog.Columns().UserId, userID).
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
