package cache

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/text/gstr"
)

type sCache struct {
	instance *gcache.Cache
}

func New() service.ICache {
	return &sCache{
		instance: g.DB().GetCache(),
	}
}

func (s *sCache) RemovePrefix(ctx context.Context, prefix string) error {
	keysToRemove := make([]any, 0, 20) //nolint:gomnd
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
	keysToRemove := make([]any, 0, 20) //nolint:gomnd
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

func (s *sCache) ClearCacheAfterPollUpdated(ctx context.Context, userID, pollID uint, sentenceUUID string) {
	if e := service.Cache().RemovePrefixes(ctx, []string{
		"poll:list",
		"poll_logs:uid:" + gconv.String(userID),
	}); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
	if e := g.DB().GetCache().Removes(ctx, g.Slice{
		"poll:id:" + gconv.String(pollID),
		"poll:sentence_uuid:" + sentenceUUID,
		"poll_log:id:" + gconv.String(pollID),
		"poll_log:sentence_uuid:" + sentenceUUID,
		"poll_marks:pid:" + gconv.String(pollID),
		"user:poll:uid:" + gconv.String(userID),
	}); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
}
