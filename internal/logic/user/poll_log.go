package user

import (
	"context"
	"fmt"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/hitokoto-osc/reviewer/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

func (s *sUser) GetUserPollLogsByUserID(ctx context.Context, userID uint) (res []entity.PollLog, err error) {
	err = dao.PollLog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 3, // 3s
		Name:     "poll_log:uid:" + gconv.String(userID),
		Force:    false,
	}).Where(do.PollLog{UserId: userID}).Scan(&res)
	return
}

func (s *sUser) GetUserPollLogs(ctx context.Context) (res []entity.PollLog, err error) {
	bizctx := service.BizCtx().Get(ctx)
	if bizctx == nil || bizctx.User == nil {
		err = gerror.New("bizctx or bizctx.User is nil")
		return
	}
	user := bizctx.User
	return s.GetUserPollLogsByUserID(ctx, user.Id)
}

func (s *sUser) GetUserPollLogsWithSentences(ctx context.Context) ([]model.PollLogWithSentence, error) {
	pollLogs, err := s.GetUserPollLogs(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]model.PollLogWithSentence, len(pollLogs))
	for i, pollLog := range pollLogs {
		var sentence *model.HitokotoV1Schema
		sentence, err = service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, pollLog.SentenceUuid)
		if err != nil {
			return nil, err
		}
		res[i] = model.PollLogWithSentence{
			PollLog:  pollLog,
			Sentence: sentence,
		}
	}
	return res, nil
}

func (s *sUser) GetUserPollLogsByUserIDWithPages(ctx context.Context, userID uint, offset, limit int) (res []entity.PollLog, err error) {
	err = dao.PollLog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 3,
		Name:     fmt.Sprintf("poll_log:uid:%d:offset:%d:limit:%d", userID, offset, limit),
		Force:    false,
	}).
		Where(do.PollLog{UserId: userID}).
		Offset(offset).
		Limit(limit).
		Scan(&res)
	return
}

func (s *sUser) GetUserPollLogsWithPages(ctx context.Context, offset, limit int) ([]entity.PollLog, error) {
	bizctx := service.BizCtx().Get(ctx)
	if bizctx == nil || bizctx.User == nil {
		err := gerror.New("bizctx or bizctx.User is nil")
		return nil, err
	}
	user := bizctx.User
	return s.GetUserPollLogsByUserIDWithPages(ctx, user.Id, offset, limit)
}

func (s *sUser) GetUserPollLogsWithSentencesAndPages(ctx context.Context, offset, limit int) ([]model.PollLogWithSentence, error) {
	pollLogs, err := s.GetUserPollLogsWithPages(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	res := make([]model.PollLogWithSentence, len(pollLogs))
	for i, pollLog := range pollLogs {
		var sentence *model.HitokotoV1Schema
		sentence, err = service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, pollLog.SentenceUuid)
		if err != nil {
			return nil, err
		}
		res[i] = model.PollLogWithSentence{
			PollLog:  pollLog,
			Sentence: sentence,
		}
	}
	return res, nil
}
