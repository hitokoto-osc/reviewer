package poll

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"golang.org/x/sync/errgroup"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) GetPollDetail(ctx context.Context, req *v1.GetPollDetailReq) (res *v1.GetPollDetailRes, err error) {
	poll, err := service.Poll().GetPollByID(ctx, req.ID)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "获取投票失败")
	} else if poll == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}
	// fetch logs and sentence
	var (
		logs     []entity.PollLog
		sentence *model.HitokotoV1Schema
	)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		logs, err = service.Poll().GetPollLogsBySentenceUUID(egCtx, poll.SentenceUuid)
		if err != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, err, "获取投票日志失败")
		}
		return nil
	})
	eg.Go(func() error {
		sentence, err = service.Hitokoto().GetHitokotoV1SchemaByUUID(egCtx, poll.SentenceUuid)
		if err != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, err, "获取句子失败")
		}
		return nil
	})
	err = eg.Wait()
	if err != nil {
		return nil, err
	}
	var records []model.PollRecord
	if len(logs) > 0 {
		records = make([]model.PollRecord, len(logs))
		for i, log := range logs {
			records[i] = model.PollRecord{
				UserID:    uint(log.UserId),
				Point:     log.Point,
				Method:    consts.PollMethod(log.Type),
				Comment:   log.Comment,
				CreatedAt: (*time.Time)(log.CreatedAt),
				UpdatedAt: (*time.Time)(log.UpdatedAt),
			}
		}
	} else {
		records = []model.PollRecord{}
	}
	res = &v1.GetPollDetailRes{
		PollElement: model.PollElement{
			SentenceUUID:       poll.SentenceUuid,
			Sentence:           sentence,
			Status:             consts.PollStatus(poll.Status),
			Approve:            poll.Accept,
			Reject:             poll.Reject,
			NeedModify:         poll.NeedEdited,
			NeedCommonUserPoll: poll.NeedUserPoll,
			CreatedAt:          (*time.Time)(poll.CreatedAt),
			UpdatedAt:          (*time.Time)(poll.UpdatedAt),
		},
		Records: records,
	}
	return res, nil
}
