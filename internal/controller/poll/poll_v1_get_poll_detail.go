package poll

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/service"

	v1 "github.com/hitokoto-osc/reviewer/api/poll/v1"
)

func (c *ControllerV1) GetPollDetail(ctx context.Context, req *v1.GetPollDetailReq) (res *v1.GetPollDetailRes, err error) {
	poll, err := service.Poll().GetPollDetailByID(ctx, uint(req.ID), req.WithPolledData)
	if err != nil {
		return nil, err
	}
	return &v1.GetPollDetailRes{PollListElement: *poll}, nil
}
