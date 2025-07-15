package wrapper

import (
	"errors"
	"github.com/gin-gonic/gin"
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

func GetCachePrefixFromContext(ctx *gin.Context) (string, error) {
	cachePrefix, ok := ctx.Get("cachePrefix")
	if !ok {
		return "", errors.New("cachePrefix not found in context")
	}
	return cachePrefix.(string), nil
}

func (cli *RedisClient) Set(ctx *gin.Context, key string, value any, expiration time.Duration) error {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return err
	}
	fullKey := cachePrefix + key
	return cli.client.Set(ctx, fullKey, value, expiration).Err()
}

func (cli *RedisClient) Del(ctx *gin.Context, keys ...string) error {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return err
	}

	prefixedKeys := make([]string, len(keys))
	for i, key := range keys {
		prefixedKeys[i] = cachePrefix + key
	}
	return cli.client.Del(ctx, prefixedKeys...).Err()
}

func (cli *RedisClient) Get(ctx *gin.Context, k string) (*redis.StringCmd, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fullKey := cachePrefix + k
	v := cli.client.Get(ctx, fullKey)
	if v.Err() != nil {
		return nil, v.Err()
	}
	return v, nil
}

func (cli *RedisClient) GetInt(ctx *gin.Context, k string) (int, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return 0, err
	}
	fullKey := cachePrefix + k
	vCmd := cli.client.Get(ctx, fullKey)
	if vCmd.Err() != nil {
		return 0, vCmd.Err()
	}
	v, err := vCmd.Int()
	return v, err
}

func (cli *RedisClient) GetString(ctx *gin.Context, k string) (string, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return "", err
	}
	fullKey := cachePrefix + k
	vCmd := cli.client.Get(ctx, fullKey)
	if vCmd.Err() != nil {
		return "", vCmd.Err()
	}
	v, err := vCmd.Result()
	return v, err
}

func (cli *RedisClient) HSet(ctx *gin.Context, key string, values ...any) (int64, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return 0, err
	}
	fullKey := cachePrefix + key
	vCmd := cli.client.HSet(ctx, fullKey, values...)
	if vCmd.Err() != nil {
		return 0, vCmd.Err()
	}
	v, err := vCmd.Result()
	return v, err
}

func (cli *RedisClient) HGetAll(ctx *gin.Context, key string) (map[string]string, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fullKey := cachePrefix + key
	v, err := cli.client.HGetAll(ctx, fullKey).Result()
	return v, err
}

func (cli *RedisClient) HDel(ctx *gin.Context, key string, fields ...string) (int64, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return 0, err
	}
	fullKey := cachePrefix + key
	v, err := cli.client.HDel(ctx, fullKey, fields...).Result()
	return v, err
}

func (cli *RedisClient) HGet(ctx *gin.Context, key, field string) (string, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return "", err
	}
	fullKey := cachePrefix + key
	v, err := cli.client.HGet(ctx, fullKey, field).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (cli *RedisClient) HMGet(ctx *gin.Context, key string, fields ...string) ([]any, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fullKey := cachePrefix + key
	v, err := cli.client.HMGet(ctx, fullKey, fields...).Result()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (cli *RedisClient) Incr(ctx *gin.Context, key string) (int64, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return 0, err
	}
	fullKey := cachePrefix + key
	v, err := cli.client.Incr(ctx, fullKey).Result()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (cli *RedisClient) Expire(ctx *gin.Context, key string, expiration time.Duration) (bool, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return false, err
	}
	fullKey := cachePrefix + key
	v, err := cli.client.Expire(ctx, fullKey, expiration).Result()
	if err != nil {
		return false, err
	}
	return v, nil
}

func (cli *RedisClient) ZRange(ctx *gin.Context, key string, start, stop int64) ([]string, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return nil, err
	}
	fullKey := cachePrefix + key
	vCmd := cli.client.ZRange(ctx, fullKey, start, stop)
	if vCmd.Err() != nil {
		return nil, vCmd.Err()
	}
	v, err := vCmd.Result()
	return v, err
}

func (cli *RedisClient) ZCard(ctx *gin.Context, key string) (int64, error) {
	cachePrefix, err := GetCachePrefixFromContext(ctx)
	if err != nil {
		return 0, err
	}
	fullKey := cachePrefix + key
	vCmd := cli.client.ZCard(ctx, fullKey)
	if vCmd.Err() != nil {
		return 0, vCmd.Err()
	}
	v, err := vCmd.Result()
	return v, err
}
