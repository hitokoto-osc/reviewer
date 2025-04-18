package poll

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"golang.org/x/sync/errgroup"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

func DailyReport(ctx context.Context) error {
	g.Log().Debug(ctx, "开始执行每日投票报告任务...")
	defer g.Log().Debug(ctx, "每日投票报告任务执行完成")
	var (
		pipelines   []entity.PollPipeline
		pollsActive []entity.Poll
		users       []entity.Users
		now         = gtime.Now()
	)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var e error
		users, e = getReviewsAndAdminsThatShouldDoNotification(egCtx)
		return gerror.Wrap(e, "获取用户列表失败")
	})
	eg.Go(func() error {
		var e error
		pipelines, e = getPollPipelinesPastDay(egCtx, now)
		return gerror.Wrap(e, "获取投票处理记录失败")
	})
	eg.Go(func() error {
		var e error
		pollsActive, e = getActivePolls(egCtx)
		return gerror.Wrap(e, "获取投票列表失败")
	})
	err := eg.Wait()
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	accept, reject, needEdited := calcProcessedFieldCount(pipelines)
	// 生成系统信息
	systemInformation := &model.DailyReportSystemInformation{
		Total:             len(pollsActive),
		ProcessTotal:      len(pipelines),
		ProcessAccept:     accept,
		ProcessReject:     reject,
		ProcessNeedEdited: needEdited,
	}
	g.Log().Debug(ctx, "开始为每个用户生成报告...")
	if len(users) == 0 {
		g.Log().Debug(ctx, "没有用户需要发送通知，因此不生成报告")
		return nil
	}
	msgChan := make(chan *model.DailyReportNotificationMessage)
	wg := sync.WaitGroup{}
	wg.Add(len(users))
	for i := 0; i < len(users); i++ {
		user := users[i]
		go func() {
			defer wg.Done()
			g.Log().Debugf(ctx, "开始为用户 %d（%s） 生成报告...", user.Id, user.Name)
			msg, e := generateDailyReportForUser(ctx, &user, systemInformation, pipelines, pollsActive, now)
			if e != nil {
				e = gerror.Wrapf(e, "生成用户 %d（%s）的报告失败", user.Id, user.Name)
				g.Log().Error(ctx, e)
				return
			}
			msgChan <- msg
		}()
	}
	go func() {
		wg.Wait()
		close(msgChan) // 关闭通道
	}()
	// 收集结果
	msgs := make([]model.DailyReportNotificationMessage, 0, len(users))
	for msg := range msgChan {
		msgs = append(msgs, *msg)
	}
	g.Log().Debug(ctx, "开始发送通知...")
	err = service.Notification().DailyReportNotification(ctx, msgs)
	if err != nil {
		return gerror.Wrap(err, "发送通知失败")
	}
	return nil
}

func generateDailyReportForUser(ctx context.Context,
	user *entity.Users,
	sysInfo *model.DailyReportSystemInformation,
	pipelines []entity.PollPipeline,
	pollInWaiting []entity.Poll,
	now *gtime.Time,
) (*model.DailyReportNotificationMessage, error) {
	pollLogs, err := getUserPollLogsPastDay(ctx, user.Id, now)
	if err != nil {
		return nil, gerror.Wrap(err, "获取用户投票记录失败")
	}
	userInfo := generateUserInformation(pipelines, pollInWaiting, pollLogs, sysInfo)
	msg := &model.DailyReportNotificationMessage{
		CreatedAt:         gtime.Now().Format("c"),
		To:                user.Email,
		UserName:          user.Name,
		SystemInformation: *sysInfo,
		UserInformation:   *userInfo,
	}
	g.Log().Debugf(ctx, "用户 %d（%s）的报告生成完成。", user.Id, user.Name)
	return msg, nil
}

// generateUserInformation 生成用户信息（通知的字段）
func generateUserInformation(
	pipelines []entity.PollPipeline,
	pollInWaiting []entity.Poll,
	pollLogs []entity.PollLog,
	sysInfo *model.DailyReportSystemInformation,
) *model.DailyReportUserInformation {
	approve, reject, needModify := calcPolledFieldCount(pollLogs)
	polled := &model.DailyReportMessageUserInformationPolled{
		Total:      len(pollLogs),
		Accept:     approve,
		Reject:     reject,
		NeedEdited: needModify,
	}
	userInfo := &model.DailyReportUserInformation{
		Polled:         *polled,
		Waiting:        0,
		Accepted:       0,
		Rejected:       0,
		InNeedEdited:   0,
		WaitForPolling: 0,
	}

	pollProcessedPastDayMap := make(g.MapIntInt)
	for _, pipeline := range pipelines {
		pollProcessedPastDayMap[pipeline.PollId] = pipeline.Operate
	}
	pollInWaitingMap := make(g.MapIntBool)
	for _, poll := range pollInWaiting {
		pollInWaitingMap[poll.Id] = true
	}
	for i := 0; i < len(pollLogs); i++ {
		log := &pollLogs[i]
		if op, ok := pollProcessedPastDayMap[log.PollId]; ok {
			switch op {
			case int(consts.PollStatusApproved):
				userInfo.Accepted++
			case int(consts.PollStatusRejected):
				userInfo.Rejected++
			case int(consts.PollStatusNeedModify):
				userInfo.InNeedEdited++
			}
		}
		if pollInWaitingMap[log.PollId] {
			userInfo.WaitForPolling++
		}
	}
	userInfo.WaitForPolling = sysInfo.Total - userInfo.Waiting // 修正需要用户投票的数据
	return userInfo
}

func getReviewsAndAdminsThatShouldDoNotification(ctx context.Context) ([]entity.Users, error) {
	var users []entity.Users
	err := dao.Users.Ctx(ctx).Raw(
		fmt.Sprintf(
			"SELECT * FROM `%s` WHERE `%s` IN (SELECT `%s` FROM `%s` WHERE `%s` = 1 AND `%s` = 1) AND (`%s` = 1 OR `%s` = 1)",
			dao.Users.Table(),
			dao.Users.Columns().Id,
			dao.UserNotification.Columns().UserId,
			dao.UserNotification.Table(),
			dao.UserNotification.Columns().EmailNotificationGlobal,
			dao.UserNotification.Columns().EmailNotificationPollDailyReport,
			dao.Users.Columns().IsReviewer,
			dao.Users.Columns().IsAdmin,
		),
	).Scan(&users)
	return users, err
}

func getPollPipelinesPastDay(ctx context.Context, now *gtime.Time) ([]entity.PollPipeline, error) {
	var pipelines []entity.PollPipeline
	err := dao.PollPipeline.Ctx(ctx).
		WhereBetween(dao.PollPipeline.Columns().CreatedAt,
			now.Add(-time.Hour*24), // nolint:mnd // 24 小时
			now,
		).Scan(&pipelines)
	return pipelines, err
}

func getUserPollLogsPastDay(ctx context.Context, userID uint, now *gtime.Time) ([]entity.PollLog, error) {
	var logs []entity.PollLog
	err := dao.PollLog.Ctx(ctx).
		Where(dao.PollLog.Columns().UserId, userID).
		WhereBetween(dao.PollLog.Columns().CreatedAt,
			now.Add(-time.Hour*24), // nolint:mnd // 24 小时
			now,
		).Scan(&logs)
	return logs, err
}

func getActivePolls(ctx context.Context) ([]entity.Poll, error) {
	var polls []entity.Poll
	err := dao.Poll.Ctx(ctx).
		WhereLT(dao.Poll.Columns().Status, int(consts.PollStatusClosed)).
		Scan(&polls)
	return polls, err
}

func calcProcessedFieldCount(pipelines []entity.PollPipeline) (approve, reject, needModify int) {
	for _, pipeline := range pipelines {
		switch pipeline.Operate {
		case int(consts.PollStatusApproved):
			approve++
		case int(consts.PollStatusRejected):
			reject++
		case int(consts.PollStatusNeedModify):
			needModify++
		}
	}
	return
}

func calcPolledFieldCount(logs []entity.PollLog) (approve, reject, needModify int) {
	for _, log := range logs {
		switch log.Type {
		case int(consts.PollMethodApprove):
			approve += 1
		case int(consts.PollMethodReject):
			reject += 1
		case int(consts.PollMethodNeedModify):
			needModify += 1
		}
	}
	return
}
