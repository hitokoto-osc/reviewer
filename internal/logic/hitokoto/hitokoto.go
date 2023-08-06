package hitokoto

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/hitokoto-osc/reviewer/internal/model"
)

type sHitokoto struct{}

func init() {
	service.RegisterHitokoto(New())
}

func New() service.IHitokoto {
	return &sHitokoto{}
}

func (s *sHitokoto) GetHitokotoV1SchemaByUUID(ctx context.Context, uuid string) (*model.HitokotoV1Schema, error) {
	var (
		hitokoto *model.HitokotoV1Schema
		err      error
	)
	// 获取顺序 Cache -> Pending -> Sentence -> Refuse
	v, err := gcache.Get(ctx, "hitokoto:uuid:"+uuid)
	if err != nil {
		return nil, err
	} else if !v.IsNil() {
		return v.Val().(*model.HitokotoV1Schema), nil
	}
	// 从 Pending 中获取
	hitokotoInPending, err := s.GetPendingByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	} else if hitokotoInPending != nil {
		hitokoto, err = s.ConvertPendingToSchemaV1(ctx, hitokotoInPending)
		if err != nil {
			return nil, err
		}
		err = gcache.Set(ctx, "hitokoto:uuid:"+uuid, hitokoto, time.Minute)
		g.Log().Error(ctx, err)
		return hitokoto, nil
	}
	// 从 Sentence 中获取
	hitokotoInSentence, err := s.GetSentenceByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if hitokotoInSentence != nil {
		hitokoto, err = s.ConvertSentenceToSchemaV1(ctx, hitokotoInSentence)
		if err != nil {
			return nil, err
		}
		err = gcache.Set(ctx, "hitokoto:uuid:"+uuid, hitokoto, time.Minute)
		g.Log().Error(ctx, err)
		return hitokoto, nil
	}
	// 从 Refuse 中获取
	hitokotoInRefuse, err := s.GetRefuseByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	} else if hitokotoInRefuse == nil {
		return nil, nil
	}
	hitokoto, err = s.ConvertRefuseToSchemaV1(ctx, hitokotoInRefuse)
	if err != nil {
		return nil, err
	}
	err = gcache.Set(ctx, "hitokoto:uuid:"+uuid, hitokoto, time.Minute)
	g.Log().Error(ctx, err)
	return hitokoto, nil
}
