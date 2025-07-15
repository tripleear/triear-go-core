package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/tripleear/triear-go-admin-core/sdk/pkg/wrapper"
	"time"
)

// NewRedis redis模式
func NewRedis(client *wrapper.RedisClient, options *redis.Options) (*Redis, error) {
	if client == nil {
		client = wrapper.NewRedisClient(options)
	}
	r := &Redis{
		client: client,
	}
	err := r.connect()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client *wrapper.RedisClient
}

func (r *Redis) GetRawRedis() *redis.Client {
	return r.client.GetRawRedis()
}

func (*Redis) String() string {
	return "redis"
}

// connect connect test
func (r *Redis) connect() error {
	var err error
	_, err = r.GetRawRedis().Ping(context.TODO()).Result()
	return err
}

// Get from key
func (r *Redis) Get(key string) (string, error) {
	return r.GetRawRedis().Get(context.TODO(), key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(key string, val interface{}, expire int) error {
	return r.GetRawRedis().Set(context.TODO(), key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(key string) error {
	return r.GetRawRedis().Del(context.TODO(), key).Err()
}

// HashGet from key
func (r *Redis) HashGet(hk, key string) (string, error) {
	return r.GetRawRedis().HGet(context.TODO(), hk, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(hk, key string) error {
	return r.GetRawRedis().HDel(context.TODO(), hk, key).Err()
}

// Increase
func (r *Redis) Increase(key string) error {
	return r.GetRawRedis().Incr(context.TODO(), key).Err()
}

func (r *Redis) Decrease(key string) error {
	return r.GetRawRedis().Decr(context.TODO(), key).Err()
}

// Set ttl
func (r *Redis) Expire(key string, dur time.Duration) error {
	return r.GetRawRedis().Expire(context.TODO(), key, dur).Err()
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *wrapper.RedisClient {
	return r.client
}
