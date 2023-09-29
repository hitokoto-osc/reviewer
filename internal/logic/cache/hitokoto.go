package cache

import (
	"context"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/samber/lo"

	"github.com/gogf/gf/v2/frame/g"
)

func (s *sCache) ClearCacheAfterHitokotoUpdated(ctx context.Context, uuids []string) {
	keys := lo.Map(uuids, func(uuid string, index int) string {
		return "uuid:" + uuid
	})
	for _, key := range keys {
		_, err := gcache.Remove(ctx, "hitokoto:"+key)
		if err != nil {
			g.Log().Error(ctx, err)
		}
		err = s.instance.Removes(ctx, []any{
			"SelectCache:pending:" + key,
			"SelectCache:refuse:" + key,
			"SelectCache:sentence:" + key,
		})
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
}
