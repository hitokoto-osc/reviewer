package job

import (
	"context"
	"sync"

	"github.com/hitokoto-osc/reviewer/internal/logic/job/poll"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

const PollTickTaskCron = "@every 1m30s"    // 每 90 秒执行一次
const PollDailyTaskCron = "0 30 8 */1 * *" // 每天八点半执行
// const PollDailyTaskCron = "@every 30s"

func DoPollTickTask(ctx context.Context) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	e := make(chan error, 2)
	go func() {
		defer wg.Done()
		err := poll.RemoveInvalidPolls(ctx)
		if err != nil {
			e <- err
		}
	}()
	go func() {
		defer wg.Done()
		err := poll.MoveOverduePolls(ctx)
		if err != nil {
			e <- err
		}
	}()
	go func() {
		wg.Wait()
		close(e)
	}()
	for err := range e {
		g.Log().Error(ctx, err)
	}
	err := poll.DoPollRuling(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func DoPollDailyTask(ctx context.Context) {
	err := poll.ClearInactiveReviewer(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	err = poll.DailyReport(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	err = poll.CalcReviewerAdoptionRate(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func RegisterPollTask(ctx context.Context) error {
	g.Log().Debug(ctx, "Registering Poll Task...")
	_, err := gcron.AddSingleton(ctx, PollTickTaskCron, DoPollTickTask)
	if err != nil {
		return err
	}
	_, err = gcron.AddSingleton(ctx, PollDailyTaskCron, DoPollDailyTask)
	return err
}
