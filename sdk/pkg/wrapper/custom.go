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

func (cli *RedisClient) HGetWithPrefixInPipeline(ctx context.Context, p redis.Pipeliner, key string, field string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return p.HGet(ctx, fullKey, field)
}

func (cli *RedisClient) HGetAllWithPrefixInPipeline(ctx context.Context, p redis.Pipeliner, key string) *redis.MapStringStringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return p.HGetAll(ctx, fullKey)
}

func (cli *RedisClient) DelWithPrefixInPipeline(ctx context.Context, p redis.Pipeliner, keys ...string) *redis.IntCmd {
	fullKeys := addPrefixToKeys(ctx, keys)
	return p.Del(ctx, fullKeys...)
}
