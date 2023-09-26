// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type (
	IPollMarks interface {
		List(ctx context.Context) ([]entity.PollMark, error)
		GetByID(ctx context.Context, id int) (*entity.PollMark, error)
		Update(ctx context.Context, mark *entity.PollMark) error
		Add(ctx context.Context, mark *do.PollMark) error
	}
)

var (
	localPollMarks IPollMarks
)

func PollMarks() IPollMarks {
	if localPollMarks == nil {
		panic("implement not found for interface IPollMarks, forgot register?")
	}
	return localPollMarks
}

func RegisterPollMarks(i IPollMarks) {
	localPollMarks = i
}
