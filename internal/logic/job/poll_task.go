package job

import (
	"context"
	"github.com/hitokoto-osc/reviewer/internal/logic/job/poll"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

const (
	PollTickTaskCron   = "0 */15 * * * *" // 每 15 分钟执行一次
	PollDailyTaskCron  = "0 30 8 */1 * *" // 每天八点半执行
	PollHourlyTaskCron = "0 0 * * * *"    // 每小时执行一次
)

// const PollDailyTaskCron = "@every 30s"

func DoPollTickTask(ctx context.Context) {
	err := poll.RemoveInvalidPolls(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	err = poll.DoPollRuling(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = poll.MoveOverduePolls(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func DoHourlyTask(ctx context.Context) {
	err := poll.CalcReviewerAdoptionRate(ctx)
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
}

func RegisterPollTask(ctx context.Context) error {
	g.Log().Debug(ctx, "Registering Poll Task...")
	_, err := gcron.AddSingleton(ctx, PollTickTaskCron, DoPollTickTask)
	if err != nil {
		return err
	}
	_, err = gcron.AddSingleton(ctx, PollHourlyTaskCron, DoHourlyTask)
	if err != nil {
		return err
	}
	_, err = gcron.AddSingleton(ctx, PollDailyTaskCron, DoPollDailyTask)
	return err
}
