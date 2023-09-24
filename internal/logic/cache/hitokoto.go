package cache

import (
	"context"

	"github.com/samber/lo"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sCache) ClearCacheAfterHitokotoUpdated(ctx context.Context, uuids []string) {
	prefixes := lo.Map(uuids, func(uuid string, index int) string {
		return "hitokoto:uuid:" + uuid
	})
	if e := s.RemovePrefixes(ctx, prefixes); e != nil {
		e = gerror.Wrap(e, "failed to remove cache: ")
		g.Log().Error(ctx, e)
	}
}
