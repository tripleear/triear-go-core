package wrapper

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(opts *redis.Options) *RedisClient {
	cli := &RedisClient{}
	cli.client = redis.NewClient(opts)
	return cli
}

func (cli *RedisClient) GetRawRedis() *redis.Client {
	return cli.client
}

func GetCachePrefixFromContext(ctx context.Context) string {
	if v, ok := ctx.Value("cachePrefix").(string); ok {
		return v
	}
	return ""
}

func (cli *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.Set(ctx, fullKey, value, expiration)
}

func (cli *RedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)

	prefixedKeys := make([]string, len(keys))
	for i, key := range keys {
		prefixedKeys[i] = cachePrefix + key
	}
	return cli.client.Del(ctx, prefixedKeys...)
}

func (cli *RedisClient) Get(ctx context.Context, k string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + k
	v := cli.client.Get(ctx, fullKey)
	return v
}

func (cli *RedisClient) HSet(ctx context.Context, key string, values ...any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	vCmd := cli.client.HSet(ctx, fullKey, values...)
	return vCmd
}

func (cli *RedisClient) HMSet(ctx context.Context, key string, values ...any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	vCmd := cli.client.HMSet(ctx, fullKey, values...)
	return vCmd
}

func (cli *RedisClient) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	newFullKey := cachePrefix + key
	vCmd := cli.client.Rename(ctx, fullKey, newFullKey)
	return vCmd
}

func (cli *RedisClient) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	v := cli.client.HGetAll(ctx, fullKey)
	return v
}

func (cli *RedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HDel(ctx, fullKey, fields...)
}

func (cli *RedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HGet(ctx, fullKey, field)
}

func (cli *RedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HMGet(ctx, fullKey, fields...)
}

func (cli *RedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.Incr(ctx, fullKey)
}

func (cli *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.Expire(ctx, fullKey, expiration)
}

func (cli *RedisClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	vCmd := cli.client.ZRange(ctx, fullKey, start, stop)
	return vCmd
}

func (cli *RedisClient) ZCard(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	vCmd := cli.client.ZCard(ctx, fullKey)
	return vCmd
}
