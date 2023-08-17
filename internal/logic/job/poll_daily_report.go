package job

import (
	"context"

	"github.com/gogf/gf/v2/os/gcron"
)

const PollDailyReportCron = "0 30 8 */1 * *" // 每天八点半执行

func DoPollDailyReport(ctx context.Context) {
}

func RegisterPollDailyReport(ctx context.Context) error {
	_, err := gcron.AddSingleton(ctx, PollDailyReportCron, DoPollDailyReport)
	return err
}
