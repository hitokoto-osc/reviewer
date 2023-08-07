// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type (
	IPoll interface {
		GetPollBySentenceUUID(ctx context.Context, uuid string) (poll *entity.Poll, err error)
		GetPollMarkLabels(ctx context.Context) ([]entity.PollMark, error)
		// GetPollMarksBySentenceUUID 获取指定投票的标签列表（不带用户信息）
		GetPollMarksBySentenceUUID(ctx context.Context, uuid string) ([]int, error)
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
