package wrapper

import (
	"context"
	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func addPrefixToKeys(ctx context.Context, keys []string) []string {
	prefix := GetCachePrefixFromContext(ctx)
	for i, k := range keys {
		keys[i] = prefix + k
	}
	return keys
}

func prefixZStoreKeys(ctx context.Context, store *redis.ZStore) {
	prefix := GetCachePrefixFromContext(ctx)
	for i, key := range store.Keys {
		store.Keys[i] = prefix + key
	}
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
