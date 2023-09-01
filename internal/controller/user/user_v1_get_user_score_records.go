package user

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/hitokoto-osc/reviewer/api/user/v1"
)

func (c *ControllerV1) GetUserScoreRecords(ctx context.Context, req *v1.GetUserScoreRecordsReq) (res *v1.GetUserScoreRecordsRes, err error) { //nolint:lll
	in := &model.GetUserScoreRecordsInput{
		Order:     dao.PollScorePipeline.Columns().CreatedAt + " " + req.Order,
		Page:      req.Page,
		PageSize:  req.PageSize,
		WithCache: true,
	}
	out, err := service.User().GetUserScoreRecords(ctx, in)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取用户积分记录失败")
	}
	res = &v1.GetUserScoreRecordsRes{
		GetUserScoreRecordsOutput: *out,
	}
	return res, nil
}
