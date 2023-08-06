package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserPollLog(ctx context.Context, req *v1.GetUserPollLogReq) (res *v1.GetUserPollLogRes, err error) {
	records, err := service.User().GetUserPollLogsWithSentencesAndPages(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取用户投票记录失败")
	}
	pollLogs := make([]model.UserPollLog, len(records))
	for i, v := range records {
		pollLogs[i] = model.UserPollLog{
			Point:        v.Point,
			SentenceUUID: v.SentenceUuid,
			Sentence:     v.Sentence,
			Method:       consts.PollMethod(v.Type),
			Comment:      v.Comment,
			CreatedAt:    (*time.Time)(v.CreatedAt),
			UpdatedAt:    (*time.Time)(v.UpdatedAt),
		}
	}
	return (*v1.GetUserPollLogRes)(&pollLogs), nil
}
