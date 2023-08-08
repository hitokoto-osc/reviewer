package poll

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"github.com/hitokoto-osc/reviewer/utility/time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) GetPollMarks(ctx context.Context, req *v1.GetPollMarksReq) (res *v1.GetPollMarksRes, err error) {
	marks, err := service.Poll().GetPollMarkLabels(ctx)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "获取投票标签失败")
	}
	pollMarks := make([]model.PollMark, len(marks))
	for i, mark := range marks {
		pollMarks[i] = model.PollMark{
			ID:        mark.Id,
			Text:      mark.Text,
			Level:     consts.PollMarkLevel(mark.Level),
			Property:  consts.PollMarkProperty(mark.Property),
			UpdatedAt: (*time.Time)(mark.UpdatedAt),
			CreatedAt: (*time.Time)(mark.CreatedAt),
		}
	}
	r := v1.GetPollMarksRes(pollMarks)
	return &r, nil
}
