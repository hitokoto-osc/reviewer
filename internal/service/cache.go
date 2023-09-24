// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICache interface {
		RemovePrefix(ctx context.Context, prefix string) error
		RemovePrefixes(ctx context.Context, prefixes []string) error
		ClearCacheAfterHitokotoUpdated(ctx context.Context, uuids []string)
		ClearCacheAfterPollUpdated(ctx context.Context, userID, pollID uint, sentenceUUID string)
		ClearPollListCache(ctx context.Context)
		ClearPollUserCache(ctx context.Context, userID uint)
	}
)

var (
	localCache ICache
)

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
