package cache

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

func (s *sCache) ClearCacheAfterPollUpdated(ctx context.Context, userID, pollID uint, sentenceUUID string) {
	if sentenceUUID == "" {
		poll, e := service.Poll().GetPollByID(ctx, int(pollID))
		if e != nil {
			g.Log().Warningf(ctx, "failed to get poll by id: %s", e.Error())
			return
		}
		sentenceUUID = poll.SentenceUuid
	}
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
		"SelectCache:poll_log:uuid:" + sentenceUUID,
		"SelectCache:poll_marks:pid:" + gconv.String(pollID),
		"SelectCache:user:poll:uid:" + gconv.String(userID),
		"SelectCache:user:poll:unreviewed:uid:" + gconv.String(userID) + ":count",
		"hitokoto:uuid:" + sentenceUUID,
	}); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
}

func (s *sCache) ClearPollListCache(ctx context.Context) {
	if e := s.RemovePrefix(ctx, "SelectCache:poll:list"); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
}

func (s *sCache) ClearPollUserCache(ctx context.Context, userID uint) {
	if e := s.RemovePrefixes(ctx, []string{
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
