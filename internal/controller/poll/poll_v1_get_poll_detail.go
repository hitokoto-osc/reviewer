package poll

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"

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
	user := service.BizCtx().GetUser(ctx)
	poll, err := service.Poll().GetPollByID(ctx, req.ID)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "获取投票失败")
	} else if poll == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}
	// fetch logs and sentence
	var (
		logs       []entity.PollLog
		sentence   *model.HitokotoV1Schema
		marks      []int
		polledData *model.PolledData
	)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var e error
		logs, e = service.Poll().GetPollLogsBySentenceUUID(egCtx, poll.SentenceUuid)
		if e != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取投票日志失败")
		}
		return nil
	})
	eg.Go(func() error {
		var e error
		sentence, e = service.Hitokoto().GetHitokotoV1SchemaByUUID(egCtx, poll.SentenceUuid)
		if e != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取句子失败")
		}
		return nil
	})
	eg.Go(func() error {
		var e error
		marks, e = service.Poll().GetPollMarksByPollID(egCtx, uint(poll.Id))
		if e != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取投票标签失败")
		}
		return nil
	})
	if req.WithPolledData {
		eg.Go(func() error {
			var e error
			polledData, e = service.User().GetUserPolledDataWithPollID(egCtx, service.BizCtx().GetUser(egCtx).Id, uint(poll.Id))
			if e != nil {
				return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取投票信息失败")
			}
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		return nil, err
	}
	var records []model.PollRecord
	if len(logs) > 0 && (user.Role == consts.UserRoleAdmin || poll.Status != int(consts.PollStatusOpen)) {
		records = make([]model.PollRecord, len(logs))
		for i, log := range logs {
			u, e := service.User().GetUserByID(ctx, uint(log.UserId))
			if e != nil {
				return nil, gerror.WrapCode(gcode.CodeOperationFailed, e, "获取用户信息失败")
			}
			records[i] = model.PollRecord{
				User: &model.UserPublicInfo{
					ID:        uint(log.UserId),
					Name:      u.Name,
					EmailHash: gmd5.MustEncryptString(u.Email),
				},
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
		PollListElement: model.PollListElement{
			PollElement: model.PollElement{
				ID:                 uint(poll.Id),
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
			Marks:      marks,
			PolledData: polledData,
			Records:    records,
		},
	}
	return res, nil
}
