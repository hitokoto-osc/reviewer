// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IJob interface {
		Register(ctx context.Context) error
	}
)

var (
	localJob IJob
)

func Job() IJob {
	if localJob == nil {
		panic("implement not found for interface IJob, forgot register?")
	}
	return localJob
}

func RegisterJob(i IJob) {
	localJob = i
}
