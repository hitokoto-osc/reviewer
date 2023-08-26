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
	ISearch interface {
		AddSentenceToSearch(ctx context.Context, schema *entity.Pending) error
	}
)

var (
	localSearch ISearch
)

func Search() ISearch {
	if localSearch == nil {
		panic("implement not found for interface ISearch, forgot register?")
	}
	return localSearch
}

func RegisterSearch(i ISearch) {
	localSearch = i
}
