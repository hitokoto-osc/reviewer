package user

import (
	"context"
	"fmt"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	vtime "github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

func (s *sUser) GetUserPollLogs(ctx context.Context, in model.UserPollLogsInput) (*model.UserPollLogsOutput, error) {
	if in.UserID == 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.New("user not found")
		}
		in.UserID = user.Id
	}
	query := dao.PollLog.Ctx(ctx)
	out := &model.UserPollLogsOutput{}
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
	in model.UserPollLogsInput,
) (*model.UserPollLogsWithSentenceOutput, error) {
	if in.UserID == 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return nil, gerror.New("user not found")
		}
		in.UserID = user.Id
	}
	query := dao.PollLog.Ctx(ctx)
	out := &model.UserPollLogsWithSentenceOutput{}
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
	for i, v := range pollLogs {
		sentence, err := service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, v.SentenceUuid)
		if err != nil {
			return nil, err
		}
		userPollLogs[i] = model.UserPollLogWithSentence{
			UserPollLog: model.UserPollLog{
				Point:        v.Point,
				SentenceUUID: v.SentenceUuid,
				Method:       consts.PollMethod(v.Type),
				Comment:      v.Comment,
				CreatedAt:    (*vtime.Time)(v.CreatedAt),
				UpdatedAt:    (*vtime.Time)(v.UpdatedAt),
			},
			Sentence: sentence,
		}
	}
	out.Collection = userPollLogs
	return out, nil
}
