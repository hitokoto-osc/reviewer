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

func init() {
	service.RegisterCache(New())
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
		"SelectCache:poll:list",
		"SelectCache:poll_logs:uid:" + gconv.String(userID),
		// "SelectCache:user:score:records:uid:" + gconv.String(userID),
	}); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
	if e := g.DB().GetCache().Removes(ctx, g.Slice{
		"SelectCache:poll:id:" + gconv.String(pollID),
		"SelectCache:poll:sentence_uuid:" + sentenceUUID,
		"SelectCache:poll_log:id:" + gconv.String(pollID),
		"SelectCache:poll_log:sentence_uuid:" + sentenceUUID,
		"SelectCache:poll_marks:pid:" + gconv.String(pollID),
		"SelectCache:user:poll:uid:" + gconv.String(userID),
		"SelectCache:user:poll:unreviewed:uid:" + gconv.String(userID) + ":count",
	}); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
}

func (s *sCache) ClearPollListCache(ctx context.Context) {
	if e := service.Cache().RemovePrefix(ctx, "SelectCache:poll:list"); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
}

func (s *sCache) ClearPollUserCache(ctx context.Context, userID uint) {
	if e := service.Cache().RemovePrefixes(ctx, []string{
		"SelectCache:poll_logs:uid:" + gconv.String(userID),
		"SelectCache:user:score:records:uid:" + gconv.String(userID),
	}); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
	if e := g.DB().GetCache().Removes(ctx, g.Slice{
		"SelectCache:user:poll:uid:" + gconv.String(userID),
		"SelectCache:user:poll:unreviewed:uid:" + gconv.String(userID) + ":count",
	}); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
}
