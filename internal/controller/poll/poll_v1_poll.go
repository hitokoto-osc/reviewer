package poll

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"golang.org/x/sync/errgroup"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) Poll(ctx context.Context, req *v1.PollReq) (res *v1.PollRes, err error) {
	if req.Method != consts.PollMethodApprove && req.Comment == "" && len(req.MarkIDs) == 0 { // 因为验证器规则太简单了，所以这里需要手动判断
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "请至少选择一个标记或者填写理由。")
	}

	var (
		poll       *entity.Poll
		polledData *model.PolledData
	)
	user := service.BizCtx().GetUser(ctx)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var e error
		poll, e = service.Poll().GetPollByID(egCtx, req.ID)
		return e
	})
	eg.Go(func() error {
		var e error
		polledData, e = service.User().GetUserPolledDataWithPollID(egCtx, user.Id, uint(req.ID))
		return e
	})
	err = eg.Wait()
	if err != nil {
		return nil, err
	} else if poll == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "投票不存在")
	}
	if polledData != nil {
		return nil, gerror.NewCode(gcode.CodeOperationFailed, "您已经投过票了，请勿重复投票！")
	}
	if poll.Status != int(consts.PollStatusOpen) {
		return nil, gerror.NewCode(gcode.CodeOperationFailed, "该投票未开放投票。")
	}

	sentence, err := service.Hitokoto().GetHitokotoV1SchemaByUUID(ctx, poll.SentenceUuid)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取句子信息失败")
	} else if sentence == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "句子不存在")
	}
	if sentence.CreatorUID == user.Id && req.Method == consts.PollMethodApprove {
		return nil, gerror.NewCode(gcode.CodeOperationFailed, "您不能给自己提交的句子投赞成票。")
	}
	points := service.Poll().GetPointsByRole(user.Role)
	err = service.Poll().Poll(ctx, &model.PollInput{
		Method:       req.Method,
		Point:        int(points),
		PollID:       uint(req.ID),
		SentenceUUID: poll.SentenceUuid,
		Comment:      req.Comment,
		UserID:       user.Id,
		IsAdmin:      user.Role == consts.UserRoleAdmin,
		Marks:        req.MarkIDs,
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "投票失败：")
	}
	return nil, nil
}
