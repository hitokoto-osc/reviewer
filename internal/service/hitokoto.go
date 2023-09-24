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
	IHitokoto interface {
		// ConvertPendingToSchemaV1 将 Pending 数据转换为 Schema V1
		ConvertPendingToSchemaV1(ctx context.Context, pending *entity.Pending) (*model.HitokotoV1WithPoll, error)
		// ConvertSentenceToSchemaV1 将 Sentence 数据转换为 Schema V1
		ConvertSentenceToSchemaV1(ctx context.Context, sentence *entity.Sentence) (*model.HitokotoV1WithPoll, error)
		// ConvertRefuseToSchemaV1 将 Refuse 数据转换为 Schema V1
		ConvertRefuseToSchemaV1(ctx context.Context, refuse *entity.Refuse) (*model.HitokotoV1WithPoll, error)
		ConvertStatusToPollStatus(status consts.HitokotoStatus) consts.PollStatus
		GetHitokotoV1SchemaByUUID(ctx context.Context, uuid string) (*model.HitokotoV1WithPoll, error)
		// GetList 获取句子列表
		//
		//nolint:gocyclo
		GetList(ctx context.Context, in *model.GetHitokotoV1SchemaListInput) (*model.GetHitokotoV1SchemaListOutput, error)
		// Move 移动句子
		//
		//nolint:gocyclo
		Move(ctx context.Context, sentence *model.HitokotoV1WithPoll, target consts.HitokotoStatus) error
		// Delete 使用 UUIDs 删除句子
		Delete(ctx context.Context, uuids []string) error
		UpdateByUUID(ctx context.Context, sentence *model.HitokotoV1WithPoll, do *model.DoHitokotoV1Update) error
		GetPendingByUUID(ctx context.Context, uuid string) (hitokoto *entity.Pending, err error)
		TopPendingPollNotOpen(ctx context.Context) (hitokoto *entity.Pending, err error)
		CountPendingPollNotOpen(ctx context.Context) (count int, err error)
		GetRefuseByUUID(ctx context.Context, uuid string) (hitokoto *entity.Refuse, err error)
		GetSentenceByUUID(ctx context.Context, uuid string) (hitokoto *entity.Sentence, err error)
	}
)

var (
	localHitokoto IHitokoto
)

func Hitokoto() IHitokoto {
	if localHitokoto == nil {
		panic("implement not found for interface IHitokoto, forgot register?")
	}
	return localHitokoto
}

func RegisterHitokoto(i IHitokoto) {
	localHitokoto = i
}
