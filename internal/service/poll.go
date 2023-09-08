// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type (
	IPoll interface {
		GetPointsByRole(role consts.UserRole) consts.UserPollPoints
		ConvertPollLogToPollRecord(ctx context.Context, in *entity.PollLog, isAdmin bool) (out *model.PollRecord, err error)
		MustConvertPollLogToPollRecord(ctx context.Context, in *entity.PollLog, isAdmin bool) (out *model.PollRecord)
		GetPollByID(ctx context.Context, pid int) (poll *entity.Poll, err error)
		// GetPollBySentenceUUID 根据 Sentence UUID 获取最新发起的投票
		GetPollBySentenceUUID(ctx context.Context, sentenceUUID string) (poll *entity.Poll, err error)
		CountOpenedPoll(ctx context.Context) (int, error)
		CreatePollByPending(ctx context.Context, pending *entity.Pending) (*entity.Poll, error)
		//nolint:gocyclo
		GetPollList(ctx context.Context, in *model.GetPollListInput) (*model.GetPollListOutput, error)
		// Poll 投票
		Poll(ctx context.Context, in *model.PollInput) error
		CancelPollByID(ctx context.Context, in *model.CancelPollInput) error
		CountUserUnreviewedPoll(ctx context.Context, uid uint) (int, error)
		GetPollLogsBySentenceUUID(ctx context.Context, uuid string) ([]entity.PollLog, error)
		GetPollLogsByPollID(ctx context.Context, pid int) ([]entity.PollLog, error)
		GetPollMarkLabels(ctx context.Context) ([]entity.PollMark, error)
		// GetPollMarksByPollID 获取指定投票的标签列表（不带用户信息）
		GetPollMarksByPollID(ctx context.Context, pid uint) ([]int, error)
		GetPollMarksByPollIDAndUserID(ctx context.Context, pid, userID int) ([]int, error)
		GetRulingThreshold(isExpandedPoll bool, totalTickets int) int
		// DoRuling 处理投票
		// nolint:gocyclo
		DoRuling(ctx context.Context, poll *entity.Poll, target consts.PollStatus) error
	}
)

var (
	localPoll IPoll
)

func Poll() IPoll {
	if localPoll == nil {
		panic("implement not found for interface IPoll, forgot register?")
	}
	return localPoll
}

func RegisterPoll(i IPoll) {
	localPoll = i
}
