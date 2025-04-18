package cache

import (
	"context"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/text/gstr"
)

type sCache struct {
	instance *gcache.Cache
}

func init() {
	service.RegisterCache(New())
}

func New() service.ICache {
	return &sCache{
		instance: g.DB().GetCache(),
	}
}

func (s *sCache) RemovePrefix(ctx context.Context, prefix string) error {
	keysToRemove := make([]any, 0, 20) //nolint:mnd
	keys, err := s.instance.KeyStrings(ctx)
	if err != nil {
		return err
	}
	for _, key := range keys {
		if gstr.HasPrefix(key, prefix) {
			keysToRemove = append(keysToRemove, key)
		}
	}
	if len(keysToRemove) > 0 {
		_, err = s.instance.Remove(ctx, keysToRemove...)
	}
	return err
}

func (s *sCache) RemovePrefixes(ctx context.Context, prefixes []string) error {
	keysToRemove := make([]any, 0, 20) //nolint:mnd
	keys, err := s.instance.KeyStrings(ctx)
	if err != nil {
		return err
	}
	for _, key := range keys {
		for _, prefix := range prefixes {
			if gstr.HasPrefix(key, prefix) {
				keysToRemove = append(keysToRemove, key)
			}
		}
	}
	if len(keysToRemove) > 0 {
		_, err = s.instance.Remove(ctx, keysToRemove...)
	}
	return err
}
