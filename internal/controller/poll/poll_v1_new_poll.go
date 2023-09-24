package poll

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/model"
	vtime "github.com/hitokoto-osc/reviewer/utility/time"

	"golang.org/x/sync/errgroup"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) NewPoll(ctx context.Context, req *v1.NewPollReq) (res *v1.NewPollRes, err error) {
	count, err := service.Poll().CountOpenedPoll(ctx)
	if err != nil {
		return nil, err
	}
	if count >= consts.PollMaxOpenPolls {
		return nil, gerror.NewCode(gcode.CodeOperationFailed, "待投票的句子过多，暂时无法创建新投票。")
	}

	var topPending *entity.Pending
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		topPending, err = service.Hitokoto().TopPendingPollNotOpen(egCtx)
		return err
	})
	eg.Go(func() error {
		count, err = service.Hitokoto().CountPendingPollNotOpen(egCtx)
		return err
	})
	if err = eg.Wait(); err != nil {
		return nil, gerror.Wrap(err, "get top pending failed")
	} else if topPending == nil {
		return nil, gerror.NewCode(gcode.CodeOperationFailed, "当前无待投票句子。")
	}
	poll, err := service.Poll().CreatePollByPending(ctx, topPending)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "创建投票失败")
	}
	res = &v1.NewPollRes{
		Poll: model.PollElement{
			SentenceUUID: poll.SentenceUuid,
			Sentence: &model.HitokotoV1WithPoll{
				HitokotoV1: model.HitokotoV1{
					ID:         uint(topPending.Id),
					UUID:       topPending.Uuid,
					Hitokoto:   topPending.Hitokoto,
					Type:       consts.HitokotoType(topPending.Type),
					From:       topPending.From,
					FromWho:    topPending.FromWho,
					Creator:    topPending.Creator,
					CreatorUID: uint(topPending.CreatorUid),
					Reviewer:   uint(topPending.Reviewer),
					CommitFrom: topPending.CommitFrom,
					Status:     consts.HitokotoStatusPending,
					CreatedAt:  topPending.CreatedAt,
				},
				PollStatus: consts.PollStatus(poll.Status),
			},
			Status:             consts.PollStatus(poll.Status),
			Approve:            poll.Accept,
			Reject:             poll.Reject,
			NeedModify:         poll.NeedEdited,
			NeedCommonUserPoll: poll.NeedUserPoll,
			CreatedAt:          (*vtime.Time)(poll.CreatedAt),
			UpdatedAt:          (*vtime.Time)(poll.UpdatedAt),
		},
		RemainPending: count - 1,
	}
	e := service.Notification().PollCreatedNotification(ctx, &res.Poll)
	if e != nil {
		g.Log().Errorf(ctx, "send poll created notification failed: %s", e.Error())
	}
	return res, nil
}
