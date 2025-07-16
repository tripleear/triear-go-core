package wrapper

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (cli *RedisClient) HSetWithPrefixInPipeline(ctx context.Context, p redis.Pipeliner, key string, values ...any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return p.HSet(ctx, fullKey, values...)
}
