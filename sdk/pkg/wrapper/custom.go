package wrapper

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (cli *RedisClient) HSetWithPrefixInPipeline(p redis.Pipeliner, ctx context.Context, key string, values ...any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return p.HSet(ctx, fullKey, values...)
}
