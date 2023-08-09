package poll

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"golang.org/x/sync/errgroup"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

// CancelPoll 目前只用于撤回普通的投票
func (c *ControllerV1) CancelPoll(ctx context.Context, req *v1.CancelPollReq) (res *v1.CancelPollRes, err error) {
	var (
		user       = service.BizCtx().GetUser(ctx)
		poll       *entity.Poll
		polledData *model.PolledData
	)
	eg, egCtx := errgroup.WithContext(ctx)
	// 获取投票信息
	eg.Go(func() (e error) {
		poll, e = service.Poll().GetPollByID(egCtx, req.ID)
		return e
	})
	// 获取已投票信息
	eg.Go(func() (e error) {
		polledData, e = service.User().GetUserPolledDataWithPollID(egCtx, user.Id, uint(req.ID))
		return e
	})
	err = eg.Wait()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "无法获取投票信息：")
	}
	if polledData == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "您尚未对此投票，无需撤回！")
	}
	if poll == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "投票不存在。可能是系統問題，建議聯係管理員。")
	}
	if poll.Status != int(consts.PollStatusOpen) {
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "此投票已結束投票階段，無法撤回投票。")
	}
	err = service.Poll().CancelPollByID(ctx, &model.CancelPollInput{
		PollID:       uint(req.ID),
		UserID:       user.Id,
		SentenceUUID: poll.SentenceUuid,
		PolledData:   polledData,
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "撤回投票失败：")
	}
	return nil, nil
}
