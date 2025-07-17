package locker

import (
	"context"
	"github.com/tripleear/triear-go-core/sdk/pkg/wrapper"
	"time"

	"github.com/bsm/redislock"
)

// NewRedis 初始化locker
func NewRedis(c *wrapper.RedisClient) *Redis {
	return &Redis{
		client: c,
	}
}

type Redis struct {
	client *wrapper.RedisClient
	mutex  *redislock.Client
}

func (r *Redis) String() string {
	return "redis"
}

func (r *Redis) Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error) {
	if r.mutex == nil {
		r.mutex = redislock.New(r.client.GetRawRedis())
	}
	return r.mutex.Obtain(context.TODO(), key, time.Duration(ttl)*time.Second, options)
}
