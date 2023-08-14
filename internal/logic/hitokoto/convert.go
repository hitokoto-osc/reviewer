package hitokoto

import (
	"context"
	"reflect"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

// convertToSchemaV1 将 Pending/Sentence/Refuse 数据转换为 Schema V1，此操作需要数据库操作查询投票状态
func (s *sHitokoto) convertToSchemaV1(ctx context.Context, data any) (*model.HitokotoV1Schema, error) {
	var (
		sentenceUUID   string
		sentenceStatus consts.HitokotoStatus
		pollStatus     consts.PollStatus
	)

	switch d := data.(type) {
	case *entity.Pending:
		if d == nil {
			err := gerror.New("pending is nil")
			return nil, err
		}
		sentenceUUID = d.Uuid
		sentenceStatus = consts.HitokotoStatusPending
	case *entity.Sentence:
		if d == nil {
			err := gerror.New("sentence is nil")
			return nil, err
		}
		sentenceUUID = d.Uuid
		sentenceStatus = consts.HitokotoStatusApproved
	case *entity.Refuse:
		if d == nil {
			err := gerror.New("refuse is nil")
			return nil, err
		}
		sentenceUUID = d.Uuid
		sentenceStatus = consts.HitokotoStatusRejected
	default:
		err := gerror.New("unknown data type")
		return nil, err
	}

	poll, err := service.Poll().GetPollBySentenceUUID(ctx, sentenceUUID)
	if err != nil {
		return nil, err
	}

	if poll == nil {
		pollStatus = consts.PollStatusNotOpen // 未开启投票
	} else {
		pollStatus = consts.PollStatus(poll.Status)
	}
	o := reflect.ValueOf(data).Elem()
	hitokoto := &model.HitokotoV1Schema{
		ID:         uint(o.FieldByName("ID").Int()),
		UUID:       o.FieldByName("Uuid").String(),
		Hitokoto:   o.FieldByName("Hitokoto").String(),
		Type:       consts.HitokotoType(o.FieldByName("Type").String()),
		From:       o.FieldByName("From").String(),
		FromWho:    o.FieldByName("FromWho").Interface().(*string),
		Creator:    o.FieldByName("Creator").String(),
		CreatorUID: uint(o.FieldByName("CreatorUid").Int()),
		Status:     sentenceStatus,
		PollStatus: pollStatus,
		CreatedAt:  o.FieldByName("CreatedAt").Interface().(string),
	}

	return hitokoto, nil
}

// ConvertPendingToSchemaV1 将 Pending 数据转换为 Schema V1
func (s *sHitokoto) ConvertPendingToSchemaV1(ctx context.Context, pending *entity.Pending) (*model.HitokotoV1Schema, error) {
	return s.convertToSchemaV1(ctx, pending)
}

// ConvertSentenceToSchemaV1 将 Sentence 数据转换为 Schema V1
func (s *sHitokoto) ConvertSentenceToSchemaV1(ctx context.Context, sentence *entity.Sentence) (*model.HitokotoV1Schema, error) {
	return s.convertToSchemaV1(ctx, sentence)
}

// ConvertRefuseToSchemaV1 将 Refuse 数据转换为 Schema V1
func (s *sHitokoto) ConvertRefuseToSchemaV1(ctx context.Context, refuse *entity.Refuse) (*model.HitokotoV1Schema, error) {
	return s.convertToSchemaV1(ctx, refuse)
}
