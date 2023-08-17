package job

import (
	"context"
	"sync"

	"github.com/hitokoto-osc/reviewer/internal/logic/job/poll"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

const PollTaskCron = "@every 1m30s" // 每 90 秒执行一次

func DoPollTask(ctx context.Context) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	e := make(chan error, 2)
	go func() {
		err := poll.RemoveInvalidPolls(ctx)
		if err != nil {
			e <- err
		}
		wg.Done()
	}()
	go func() {
		err := poll.MoveOverduePolls(ctx)
		if err != nil {
			e <- err
		}
		wg.Done()
	}()
	wg.Wait()
	if len(e) > 0 {
		for i := 0; i < len(e); i++ {
			g.Log().Error(ctx, <-e)
		}
	}
	err := poll.DoPollRuling(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func RegisterPollTask(ctx context.Context) error {
	g.Log().Debug(ctx, "Registering Poll Task...")
	_, err := gcron.AddSingleton(ctx, PollTaskCron, DoPollTask)
	return err
}
