package job

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/service"
)

type sJob struct{}

func init() {
	service.RegisterJob(New())
}

func New() service.IJob {
	return &sJob{}
}

func (s *sJob) Register(ctx context.Context) error {
	if e := RegisterPollTask(ctx); e != nil {
		return e
	}
	return nil
}
