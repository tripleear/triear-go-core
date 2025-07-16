package wrapper

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	"time"
)

func (cli *RedisClient) Pipeline() redis.Pipeliner {
	return cli.client.Pipeline()
}

func (cli *RedisClient) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return cli.client.Pipelined(ctx, fn)
}

func (cli *RedisClient) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return cli.client.TxPipelined(ctx, fn)
}

func (cli *RedisClient) TxPipeline() redis.Pipeliner {
	return cli.client.TxPipeline()
}

func (cli *RedisClient) Command(ctx context.Context) *redis.CommandsInfoCmd {
	return cli.client.Command(ctx)
}

func (cli *RedisClient) ClientGetName(ctx context.Context) *redis.StringCmd {
	return cli.client.ClientGetName(ctx)
}

func (cli *RedisClient) Echo(ctx context.Context, message any) *redis.StringCmd {
	return cli.client.Echo(ctx, message)
}

func (cli *RedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	return cli.client.Ping(ctx)
}

func (cli *RedisClient) Quit(ctx context.Context) *redis.StatusCmd {
	return cli.client.Quit(ctx)
}

func (cli *RedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)

	prefixedKeys := make([]string, len(keys))
	for i, key := range keys {
		prefixedKeys[i] = cachePrefix + key
	}
	return cli.client.Del(ctx, prefixedKeys...)
}

func (cli *RedisClient) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)

	prefixedKeys := make([]string, len(keys))
	for i, key := range keys {
		prefixedKeys[i] = cachePrefix + key
	}
	return cli.client.Unlink(ctx, prefixedKeys...)
}

func (cli *RedisClient) Dump(ctx context.Context, key string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.Dump(ctx, fullKey)
}

func (cli *RedisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)

	prefixedKeys := make([]string, len(keys))
	for i, key := range keys {
		prefixedKeys[i] = cachePrefix + key
	}
	return cli.client.Exists(ctx, prefixedKeys...)
}

func (cli *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.Expire(ctx, fullKey, expiration)
}
func (cli *RedisClient) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ExpireAt(ctx, prefix+key, tm)
}

func (cli *RedisClient) ExpireNX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ExpireNX(ctx, prefix+key, expiration)
}

func (cli *RedisClient) ExpireXX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ExpireXX(ctx, prefix+key, expiration)
}

func (cli *RedisClient) ExpireGT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ExpireGT(ctx, prefix+key, expiration)
}

func (cli *RedisClient) ExpireLT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ExpireLT(ctx, prefix+key, expiration)
}

func (cli *RedisClient) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Keys(ctx, prefix+pattern)
}

func (cli *RedisClient) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Migrate(ctx, host, port, prefix+key, db, timeout)
}

func (cli *RedisClient) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Move(ctx, prefix+key, db)
}

func (cli *RedisClient) LMove(ctx context.Context, source, destination, srcpos, destpos string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	sourceFullKey := cachePrefix + source
	destFullKey := cachePrefix + destination
	return cli.client.LMove(ctx, sourceFullKey, destFullKey, srcpos, destpos)
}

func (cli *RedisClient) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	sourceFullKey := cachePrefix + source
	destFullKey := cachePrefix + destination
	return cli.client.BLMove(ctx, sourceFullKey, destFullKey, srcpos, destpos, timeout)
}

func (cli *RedisClient) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ObjectRefCount(ctx, prefix+key)
}

func (cli *RedisClient) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ObjectEncoding(ctx, prefix+key)
}

func (cli *RedisClient) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.ObjectIdleTime(ctx, prefix+key)
}

func (cli *RedisClient) Persist(ctx context.Context, key string) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Persist(ctx, prefix+key)
}

func (cli *RedisClient) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.PExpire(ctx, prefix+key, expiration)
}

func (cli *RedisClient) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.PExpireAt(ctx, prefix+key, tm)
}

func (cli *RedisClient) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.PTTL(ctx, prefix+key)
}

func (cli *RedisClient) RandomKey(ctx context.Context) *redis.StringCmd {
	// RandomKey 无法控制前缀，它直接返回 redis 中任意 key
	// 如果你需要限制前缀，建议用 Keys 命令代替
	return cli.client.RandomKey(ctx)
}

func (cli *RedisClient) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	newFullKey := cachePrefix + newkey
	vCmd := cli.client.Rename(ctx, fullKey, newFullKey)
	return vCmd
}
func (cli *RedisClient) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.RenameNX(ctx, prefix+key, prefix+newkey)
}

func (cli *RedisClient) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Restore(ctx, prefix+key, ttl, value)
}

func (cli *RedisClient) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.RestoreReplace(ctx, prefix+key, ttl, value)
}

func (cli *RedisClient) Sort(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Sort(ctx, prefix+key, sort)
}

func (cli *RedisClient) SortStore(ctx context.Context, key, store string, sort *redis.Sort) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.SortStore(ctx, prefix+key, prefix+store, sort)
}

func (cli *RedisClient) SortInterfaces(ctx context.Context, key string, sort *redis.Sort) *redis.SliceCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.SortInterfaces(ctx, prefix+key, sort)
}

func (cli *RedisClient) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	prefixedKeys := make([]string, len(keys))
	for i, k := range keys {
		prefixedKeys[i] = prefix + k
	}
	return cli.client.Touch(ctx, prefixedKeys...)
}

func (cli *RedisClient) TTL(ctx context.Context, key string) *redis.DurationCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.TTL(ctx, prefix+key)
}

func (cli *RedisClient) Type(ctx context.Context, key string) *redis.StatusCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Type(ctx, prefix+key)
}

func (cli *RedisClient) Append(ctx context.Context, key, value string) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Append(ctx, prefix+key, value)
}

func (cli *RedisClient) Decr(ctx context.Context, key string) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Decr(ctx, prefix+key)
}

func (cli *RedisClient) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.DecrBy(ctx, prefix+key, decrement)
}

func (cli *RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.Get(ctx, prefix+key)
}

func (cli *RedisClient) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.GetRange(ctx, prefix+key, start, end)
}

func (cli *RedisClient) GetSet(ctx context.Context, key string, value any) *redis.StringCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.GetSet(ctx, prefix+key, value)
}

func (cli *RedisClient) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	prefix := GetCachePrefixFromContext(ctx)
	return cli.client.GetEx(ctx, prefix+key, expiration)
}

func (cli *RedisClient) GetDel(ctx context.Context, key string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.GetDel(ctx, fullKey)
}

func (cli *RedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.Incr(ctx, fullKey)
}

func (cli *RedisClient) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.IncrBy(ctx, fullKey, value)
}
func (cli *RedisClient) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.MGet(ctx, keys...)
}

func (cli *RedisClient) MSet(ctx context.Context, values ...any) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	prefixedValues := make([]any, len(values))
	for i := 0; i < len(values); i += 2 {
		if key, ok := values[i].(string); ok {
			prefixedValues[i] = cachePrefix + key
			prefixedValues[i+1] = values[i+1]
		}
	}
	return cli.client.MSet(ctx, prefixedValues...)
}

func (cli *RedisClient) MSetNX(ctx context.Context, values ...any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	prefixedValues := make([]any, len(values))
	for i := 0; i < len(values); i += 2 {
		if key, ok := values[i].(string); ok {
			prefixedValues[i] = cachePrefix + key
			prefixedValues[i+1] = values[i+1]
		}
	}
	return cli.client.MSetNX(ctx, prefixedValues...)
}

func (cli *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.Set(ctx, fullKey, value, expiration)
}

func (cli *RedisClient) SetArgs(ctx context.Context, key string, value any, a redis.SetArgs) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SetArgs(ctx, fullKey, value, a)
}

func (cli *RedisClient) SetEx(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SetEx(ctx, fullKey, value, expiration)
}

func (cli *RedisClient) SetNX(ctx context.Context, key string, value any, expiration time.Duration) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SetNX(ctx, fullKey, value, expiration)
}

func (cli *RedisClient) SetXX(ctx context.Context, key string, value any, expiration time.Duration) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SetXX(ctx, fullKey, value, expiration)
}

func (cli *RedisClient) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SetRange(ctx, fullKey, offset, value)
}

func (cli *RedisClient) StrLen(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.StrLen(ctx, fullKey)
}

func (cli *RedisClient) Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	sourceFullKey := cachePrefix + sourceKey
	destFullKey := cachePrefix + destKey
	return cli.client.Copy(ctx, sourceFullKey, destFullKey, db, replace)
}

func (cli *RedisClient) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.GetBit(ctx, fullKey, offset)
}

func (cli *RedisClient) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SetBit(ctx, fullKey, offset, value)
}

func (cli *RedisClient) BitCount(ctx context.Context, key string, bitCount *redis.BitCount) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BitCount(ctx, fullKey, bitCount)
}

func (cli *RedisClient) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destFullKey := cachePrefix + destKey
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.BitOpAnd(ctx, destFullKey, keys...)
}

func (cli *RedisClient) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destFullKey := cachePrefix + destKey
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.BitOpOr(ctx, destFullKey, keys...)
}

func (cli *RedisClient) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destFullKey := cachePrefix + destKey
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.BitOpXor(ctx, destFullKey, keys...)
}

func (cli *RedisClient) BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destFullKey := cachePrefix + destKey
	fullKey := cachePrefix + key
	return cli.client.BitOpNot(ctx, destFullKey, fullKey)
}

func (cli *RedisClient) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BitPos(ctx, fullKey, bit, pos...)
}

func (cli *RedisClient) BitField(ctx context.Context, key string, args ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BitField(ctx, fullKey, args...)
}

func (cli *RedisClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	return cli.client.Scan(ctx, cursor, match, count)
}

func (cli *RedisClient) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	return cli.client.ScanType(ctx, cursor, match, count, keyType)
}

func (cli *RedisClient) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SScan(ctx, fullKey, cursor, match, count)
}

func (cli *RedisClient) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HScan(ctx, fullKey, cursor, match, count)
}

func (cli *RedisClient) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZScan(ctx, fullKey, cursor, match, count)
}

func (cli *RedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HDel(ctx, fullKey, fields...)
}

func (cli *RedisClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HExists(ctx, fullKey, field)
}

func (cli *RedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HGet(ctx, fullKey, field)
}

func (cli *RedisClient) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HGetAll(ctx, fullKey)
}

func (cli *RedisClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HIncrBy(ctx, fullKey, field, incr)
}

func (cli *RedisClient) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.IncrByFloat(ctx, fullKey, value)
}

func (cli *RedisClient) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HIncrByFloat(ctx, fullKey, field, incr)
}

func (cli *RedisClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HKeys(ctx, fullKey)
}

func (cli *RedisClient) HLen(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HLen(ctx, fullKey)
}

func (cli *RedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HMGet(ctx, fullKey, fields...)
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
func (cli *RedisClient) HSetNX(ctx context.Context, key, field string, value any) *redis.BoolCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.HSetNX(ctx, fullKey, field, value)
}

func (cli *RedisClient) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.HVals(ctx, fullKey)
}

func (cli *RedisClient) HRandField(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.HRandField(ctx, fullKey, count)
}

func (cli *RedisClient) HRandFieldWithValues(ctx context.Context, key string, count int) *redis.KeyValueSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.HRandFieldWithValues(ctx, fullKey, count)
}

func (cli *RedisClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.BLPop(ctx, timeout, keys...)
}

func (cli *RedisClient) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.BRPop(ctx, timeout, keys...)
}

func (cli *RedisClient) RPush(ctx context.Context, key string, values ...any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.RPush(ctx, fullKey, values...)
}

func (cli *RedisClient) RPushX(ctx context.Context, key string, values ...any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.RPushX(ctx, fullKey, values...)
}

func (cli *RedisClient) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	return cli.client.BRPopLPush(ctx, source, destination, timeout)
}

func (cli *RedisClient) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	return cli.client.RPopLPush(ctx, source, destination)
}

func (cli *RedisClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LIndex(ctx, fullKey, index)
}

func (cli *RedisClient) LInsert(ctx context.Context, key, op string, pivot, value any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LInsert(ctx, fullKey, op, pivot, value)
}

func (cli *RedisClient) LInsertBefore(ctx context.Context, key string, pivot, value any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LInsertBefore(ctx, fullKey, pivot, value)
}

func (cli *RedisClient) LInsertAfter(ctx context.Context, key string, pivot, value any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LInsertAfter(ctx, fullKey, pivot, value)
}

func (cli *RedisClient) LLen(ctx context.Context, key string) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LLen(ctx, fullKey)
}

func (cli *RedisClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LPop(ctx, fullKey)
}

func (cli *RedisClient) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LPopCount(ctx, fullKey, count)
}

func (cli *RedisClient) LPos(ctx context.Context, key string, value string, args redis.LPosArgs) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LPos(ctx, fullKey, value, args)
}

func (cli *RedisClient) LPosCount(ctx context.Context, key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LPosCount(ctx, fullKey, value, count, args)
}

func (cli *RedisClient) LPush(ctx context.Context, key string, values ...any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LPush(ctx, fullKey, values...)
}

func (cli *RedisClient) LPushX(ctx context.Context, key string, values ...any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LPushX(ctx, fullKey, values...)
}

func (cli *RedisClient) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LRange(ctx, fullKey, start, stop)
}

func (cli *RedisClient) LRem(ctx context.Context, key string, count int64, value any) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LRem(ctx, fullKey, count, value)
}

func (cli *RedisClient) LSet(ctx context.Context, key string, index int64, value any) *redis.StatusCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LSet(ctx, fullKey, index, value)
}

func (cli *RedisClient) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.LTrim(ctx, fullKey, start, stop)
}

func (cli *RedisClient) RPop(ctx context.Context, key string) *redis.StringCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.RPop(ctx, fullKey)
}

func (cli *RedisClient) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.RPopCount(ctx, fullKey, count)
}

func (cli *RedisClient) SAdd(ctx context.Context, key string, members ...any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SAdd(ctx, fullKey, members...)
}

func (cli *RedisClient) SCard(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SCard(ctx, fullKey)
}

func (cli *RedisClient) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, k := range keys {
		keys[i] = cachePrefix + k
	}
	return cli.client.SDiff(ctx, keys...)
}

func (cli *RedisClient) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destKey := cachePrefix + destination
	for i, k := range keys {
		keys[i] = cachePrefix + k
	}
	return cli.client.SDiffStore(ctx, destKey, keys...)
}

func (cli *RedisClient) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, k := range keys {
		keys[i] = cachePrefix + k
	}
	return cli.client.SInter(ctx, keys...)
}

func (cli *RedisClient) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destKey := cachePrefix + destination
	for i, k := range keys {
		keys[i] = cachePrefix + k
	}
	return cli.client.SInterStore(ctx, destKey, keys...)
}

func (cli *RedisClient) SIsMember(ctx context.Context, key string, member any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SIsMember(ctx, fullKey, member)
}

func (cli *RedisClient) SMIsMember(ctx context.Context, key string, members ...any) *redis.BoolSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SMIsMember(ctx, fullKey, members...)
}

func (cli *RedisClient) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SMembers(ctx, fullKey)
}

func (cli *RedisClient) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SMembersMap(ctx, fullKey)
}

func (cli *RedisClient) SMove(ctx context.Context, source, destination string, member any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	sourceKey := cachePrefix + source
	destKey := cachePrefix + destination
	return cli.client.SMove(ctx, sourceKey, destKey, member)
}

func (cli *RedisClient) SPop(ctx context.Context, key string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SPop(ctx, fullKey)
}

func (cli *RedisClient) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SPopN(ctx, fullKey, count)
}

func (cli *RedisClient) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SRandMember(ctx, fullKey)
}

func (cli *RedisClient) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SRandMemberN(ctx, fullKey, count)
}

func (cli *RedisClient) SRem(ctx context.Context, key string, members ...any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.SRem(ctx, fullKey, members...)
}

func (cli *RedisClient) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, k := range keys {
		keys[i] = cachePrefix + k
	}
	return cli.client.SUnion(ctx, keys...)
}

func (cli *RedisClient) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destKey := cachePrefix + destination
	for i, k := range keys {
		keys[i] = cachePrefix + k
	}
	return cli.client.SUnionStore(ctx, destKey, keys...)
}

func (cli *RedisClient) XAdd(ctx context.Context, a *redis.XAddArgs) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	a.Stream = cachePrefix + a.Stream
	return cli.client.XAdd(ctx, a)
}

func (cli *RedisClient) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XDel(ctx, fullStream, ids...)
}

func (cli *RedisClient) XLen(ctx context.Context, stream string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XLen(ctx, fullStream)
}

func (cli *RedisClient) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XRange(ctx, fullStream, start, stop)
}

func (cli *RedisClient) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XRangeN(ctx, fullStream, start, stop, count)
}

func (cli *RedisClient) XRevRange(ctx context.Context, stream string, start, stop string) *redis.XMessageSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XRevRange(ctx, fullStream, start, stop)
}

func (cli *RedisClient) XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *redis.XMessageSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XRevRangeN(ctx, fullStream, start, stop, count)
}

func (cli *RedisClient) XRead(ctx context.Context, a *redis.XReadArgs) *redis.XStreamSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, stream := range a.Streams {
		a.Streams[i] = cachePrefix + stream
	}
	return cli.client.XRead(ctx, a)
}

func (cli *RedisClient) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, stream := range streams {
		streams[i] = cachePrefix + stream
	}
	return cli.client.XReadStreams(ctx, streams...)
}

func (cli *RedisClient) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XGroupCreate(ctx, fullStream, group, start)
}

func (cli *RedisClient) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XGroupCreateMkStream(ctx, fullStream, group, start)
}

func (cli *RedisClient) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XGroupSetID(ctx, fullStream, group, start)
}

func (cli *RedisClient) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XGroupDestroy(ctx, fullStream, group)
}

func (cli *RedisClient) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XGroupCreateConsumer(ctx, fullStream, group, consumer)
}

func (cli *RedisClient) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XGroupDelConsumer(ctx, fullStream, group, consumer)
}

func (cli *RedisClient) XReadGroup(ctx context.Context, a *redis.XReadGroupArgs) *redis.XStreamSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, stream := range a.Streams {
		a.Streams[i] = cachePrefix + stream
	}
	return cli.client.XReadGroup(ctx, a)
}

func (cli *RedisClient) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XAck(ctx, fullStream, group, ids...)
}

func (cli *RedisClient) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullStream := cachePrefix + stream
	return cli.client.XPending(ctx, fullStream, group)
}

func (cli *RedisClient) XPendingExt(ctx context.Context, a *redis.XPendingExtArgs) *redis.XPendingExtCmd {
	return cli.client.XPendingExt(ctx, a)
}

func (cli *RedisClient) XClaim(ctx context.Context, a *redis.XClaimArgs) *redis.XMessageSliceCmd {
	return cli.client.XClaim(ctx, a)
}

func (cli *RedisClient) XClaimJustID(ctx context.Context, a *redis.XClaimArgs) *redis.StringSliceCmd {
	return cli.client.XClaimJustID(ctx, a)
}

func (cli *RedisClient) XAutoClaim(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimCmd {
	return cli.client.XAutoClaim(ctx, a)
}

func (cli *RedisClient) XAutoClaimJustID(ctx context.Context, a *redis.XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	return cli.client.XAutoClaimJustID(ctx, a)
}

func (cli *RedisClient) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XTrimMaxLen(ctx, fullKey, maxLen)
}

func (cli *RedisClient) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XTrimMaxLenApprox(ctx, fullKey, maxLen, limit)
}

func (cli *RedisClient) XTrimMinID(ctx context.Context, key string, minID string) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XTrimMinID(ctx, fullKey, minID)
}

func (cli *RedisClient) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XTrimMinIDApprox(ctx, fullKey, minID, limit)
}

func (cli *RedisClient) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XInfoGroups(ctx, fullKey)
}

func (cli *RedisClient) XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XInfoStream(ctx, fullKey)
}

func (cli *RedisClient) XInfoStreamFull(ctx context.Context, key string, count int) *redis.XInfoStreamFullCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XInfoStreamFull(ctx, fullKey, count)
}

func (cli *RedisClient) XInfoConsumers(ctx context.Context, key string, group string) *redis.XInfoConsumersCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.XInfoConsumers(ctx, fullKey, group)
}

func (cli *RedisClient) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	fullKeys := addPrefixToKeys(ctx, keys)
	return cli.client.BZPopMax(ctx, timeout, fullKeys...)
}

func (cli *RedisClient) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	fullKeys := addPrefixToKeys(ctx, keys)
	return cli.client.BZPopMin(ctx, timeout, fullKeys...)
}

func (cli *RedisClient) ZAdd(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZAdd(ctx, fullKey, members...)
}

func (cli *RedisClient) ZAddNX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZAddNX(ctx, fullKey, members...)
}

func (cli *RedisClient) ZAddXX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZAddXX(ctx, fullKey, members...)
}

func (cli *RedisClient) ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZAddArgs(ctx, fullKey, args)
}

func (cli *RedisClient) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZAddArgsIncr(ctx, fullKey, args)
}

func (cli *RedisClient) ZCard(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	vCmd := cli.client.ZCard(ctx, fullKey)
	return vCmd
}

func (cli *RedisClient) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZCount(ctx, fullKey, min, max)
}

func (cli *RedisClient) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZLexCount(ctx, fullKey, min, max)
}

func (cli *RedisClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZIncrBy(ctx, fullKey, increment, member)
}

func (cli *RedisClient) ZInter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd {
	prefixZStoreKeys(ctx, store)
	return cli.client.ZInter(ctx, store)
}

func (cli *RedisClient) ZInterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd {
	prefixZStoreKeys(ctx, store)
	return cli.client.ZInterWithScores(ctx, store)
}

func (cli *RedisClient) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	prefixZStoreKeys(ctx, store)
	return cli.client.ZInterStore(ctx, prefix+destination, store)
}

func (cli *RedisClient) ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZMScore(ctx, fullKey, members...)
}

func (cli *RedisClient) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZPopMax(ctx, fullKey, count...)
}

func (cli *RedisClient) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	fullKey := GetCachePrefixFromContext(ctx) + key
	return cli.client.ZPopMin(ctx, fullKey, count...)
}

func (cli *RedisClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	vCmd := cli.client.ZRange(ctx, fullKey, start, stop)
	return vCmd
}

func (cli *RedisClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRangeWithScores(ctx, fullKey, start, stop)
}

func (cli *RedisClient) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRangeByScore(ctx, fullKey, opt)
}

func (cli *RedisClient) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRangeByLex(ctx, fullKey, opt)
}

func (cli *RedisClient) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRangeByScoreWithScores(ctx, fullKey, opt)
}

func (cli *RedisClient) ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	z.Key = cachePrefix + z.Key
	return cli.client.ZRangeArgs(ctx, z)
}

func (cli *RedisClient) ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	z.Key = cachePrefix + z.Key
	return cli.client.ZRangeArgsWithScores(ctx, z)
}

func (cli *RedisClient) ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	dstFull := cachePrefix + dst
	z.Key = cachePrefix + z.Key
	return cli.client.ZRangeStore(ctx, dstFull, z)
}

func (cli *RedisClient) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRank(ctx, fullKey, member)
}

func (cli *RedisClient) ZRem(ctx context.Context, key string, members ...any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRem(ctx, fullKey, members...)
}

func (cli *RedisClient) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRemRangeByRank(ctx, fullKey, start, stop)
}

func (cli *RedisClient) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRemRangeByScore(ctx, fullKey, min, max)
}

func (cli *RedisClient) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRemRangeByLex(ctx, fullKey, min, max)
}

func (cli *RedisClient) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRevRange(ctx, fullKey, start, stop)
}

func (cli *RedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRevRangeWithScores(ctx, fullKey, start, stop)
}

func (cli *RedisClient) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRevRangeByScore(ctx, fullKey, opt)
}

func (cli *RedisClient) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRevRangeByLex(ctx, fullKey, opt)
}

func (cli *RedisClient) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRevRangeByScoreWithScores(ctx, fullKey, opt)
}

func (cli *RedisClient) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRevRank(ctx, fullKey, member)
}

func (cli *RedisClient) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZScore(ctx, fullKey, member)
}

func (cli *RedisClient) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destFull := cachePrefix + dest
	for i, key := range store.Keys {
		store.Keys[i] = cachePrefix + key
	}
	return cli.client.ZUnionStore(ctx, destFull, store)
}

func (cli *RedisClient) ZRandMember(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRandMember(ctx, fullKey, count)
}

func (cli *RedisClient) ZRandMemberWithScores(ctx context.Context, key string, count int) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRandMemberWithScores(ctx, fullKey, count)
}

func (cli *RedisClient) ZUnion(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range store.Keys {
		store.Keys[i] = cachePrefix + key
	}
	return cli.client.ZUnion(ctx, store)
}

func (cli *RedisClient) ZUnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range store.Keys {
		store.Keys[i] = cachePrefix + key
	}
	return cli.client.ZUnionWithScores(ctx, store)
}

func (cli *RedisClient) ZDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.ZDiff(ctx, keys...)
}

func (cli *RedisClient) ZDiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.ZDiffWithScores(ctx, keys...)
}

func (cli *RedisClient) ZDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	destFull := cachePrefix + destination
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.ZDiffStore(ctx, destFull, keys...)
}

func (cli *RedisClient) PFAdd(ctx context.Context, key string, els ...any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.PFAdd(ctx, fullKey, els...)
}

func (cli *RedisClient) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	for i, key := range keys {
		keys[i] = cachePrefix + key
	}
	return cli.client.PFCount(ctx, keys...)
}
func (cli *RedisClient) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	prefix := GetCachePrefixFromContext(ctx)
	dest = prefix + dest
	for i := range keys {
		keys[i] = prefix + keys[i]
	}
	return cli.client.PFMerge(ctx, dest, keys...)
}

func (cli *RedisClient) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	return cli.client.BgRewriteAOF(ctx)
}

func (cli *RedisClient) BgSave(ctx context.Context) *redis.StatusCmd {
	return cli.client.BgSave(ctx)
}

func (cli *RedisClient) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	return cli.client.ClientKill(ctx, ipPort)
}

func (cli *RedisClient) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	return cli.client.ClientKillByFilter(ctx, keys...)
}

func (cli *RedisClient) ClientList(ctx context.Context) *redis.StringCmd {
	return cli.client.ClientList(ctx)
}

func (cli *RedisClient) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	return cli.client.ClientPause(ctx, dur)
}

func (cli *RedisClient) ClientUnpause(ctx context.Context) *redis.BoolCmd {
	return cli.client.ClientUnpause(ctx)
}

func (cli *RedisClient) ClientID(ctx context.Context) *redis.IntCmd {
	return cli.client.ClientID(ctx)
}

func (cli *RedisClient) ClientUnblock(ctx context.Context, id int64) *redis.IntCmd {
	return cli.client.ClientUnblock(ctx, id)
}

func (cli *RedisClient) ClientUnblockWithError(ctx context.Context, id int64) *redis.IntCmd {
	return cli.client.ClientUnblockWithError(ctx, id)
}

func (cli *RedisClient) ConfigGet(ctx context.Context, parameter string) *redis.MapStringStringCmd {
	return cli.client.ConfigGet(ctx, parameter)
}

func (cli *RedisClient) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	return cli.client.ConfigResetStat(ctx)
}

func (cli *RedisClient) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	return cli.client.ConfigSet(ctx, parameter, value)
}

func (cli *RedisClient) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	return cli.client.ConfigRewrite(ctx)
}

func (cli *RedisClient) DBSize(ctx context.Context) *redis.IntCmd {
	return cli.client.DBSize(ctx)
}

func (cli *RedisClient) FlushAll(ctx context.Context) *redis.StatusCmd {
	return cli.client.FlushAll(ctx)
}

func (cli *RedisClient) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	return cli.client.FlushAllAsync(ctx)
}

func (cli *RedisClient) FlushDB(ctx context.Context) *redis.StatusCmd {
	return cli.client.FlushDB(ctx)
}

func (cli *RedisClient) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	return cli.client.FlushDBAsync(ctx)
}

func (cli *RedisClient) Info(ctx context.Context, section ...string) *redis.StringCmd {
	return cli.client.Info(ctx, section...)
}

func (cli *RedisClient) LastSave(ctx context.Context) *redis.IntCmd {
	return cli.client.LastSave(ctx)
}

func (cli *RedisClient) Save(ctx context.Context) *redis.StatusCmd {
	return cli.client.Save(ctx)
}

func (cli *RedisClient) Shutdown(ctx context.Context) *redis.StatusCmd {
	return cli.client.Shutdown(ctx)
}

func (cli *RedisClient) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	return cli.client.ShutdownSave(ctx)
}

func (cli *RedisClient) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	return cli.client.ShutdownNoSave(ctx)
}

func (cli *RedisClient) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	return cli.client.SlaveOf(ctx, host, port)
}

func (cli *RedisClient) SlowLogGet(ctx context.Context, num int64) *redis.SlowLogCmd {
	return cli.client.SlowLogGet(ctx, num)
}

func (cli *RedisClient) Time(ctx context.Context) *redis.TimeCmd {
	return cli.client.Time(ctx)
}

func (cli *RedisClient) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	prefix := GetCachePrefixFromContext(ctx)
	key = prefix + key
	return cli.client.DebugObject(ctx, key)
}

func (cli *RedisClient) ReadOnly(ctx context.Context) *redis.StatusCmd {
	return cli.client.ReadOnly(ctx)
}

func (cli *RedisClient) ReadWrite(ctx context.Context) *redis.StatusCmd {
	return cli.client.ReadWrite(ctx)
}

func (cli *RedisClient) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	prefix := GetCachePrefixFromContext(ctx)
	key = prefix + key
	return cli.client.MemoryUsage(ctx, key, samples...)
}

func (cli *RedisClient) Eval(ctx context.Context, script string, keys []string, args ...any) *redis.Cmd {
	prefix := GetCachePrefixFromContext(ctx)
	for i := range keys {
		keys[i] = prefix + keys[i]
	}
	return cli.client.Eval(ctx, script, keys, args...)
}

func (cli *RedisClient) EvalSha(ctx context.Context, sha1 string, keys []string, args ...any) *redis.Cmd {
	prefix := GetCachePrefixFromContext(ctx)
	for i := range keys {
		keys[i] = prefix + keys[i]
	}
	return cli.client.EvalSha(ctx, sha1, keys, args...)
}

func (cli *RedisClient) EvalRO(ctx context.Context, script string, keys []string, args ...any) *redis.Cmd {
	prefix := GetCachePrefixFromContext(ctx)
	for i := range keys {
		keys[i] = prefix + keys[i]
	}
	return cli.client.EvalRO(ctx, script, keys, args...)
}

func (cli *RedisClient) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...any) *redis.Cmd {
	prefix := GetCachePrefixFromContext(ctx)
	for i := range keys {
		keys[i] = prefix + keys[i]
	}
	return cli.client.EvalShaRO(ctx, sha1, keys, args...)
}

func (cli *RedisClient) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return cli.client.ScriptExists(ctx, hashes...)
}

func (cli *RedisClient) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	return cli.client.ScriptFlush(ctx)
}

func (cli *RedisClient) ScriptKill(ctx context.Context) *redis.StatusCmd {
	return cli.client.ScriptKill(ctx)
}

func (cli *RedisClient) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return cli.client.ScriptLoad(ctx, script)
}

func (cli *RedisClient) Publish(ctx context.Context, channel string, message any) *redis.IntCmd {
	return cli.client.Publish(ctx, channel, message)
}

func (cli *RedisClient) SPublish(ctx context.Context, channel string, message any) *redis.IntCmd {
	return cli.client.SPublish(ctx, channel, message)
}

func (cli *RedisClient) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return cli.client.PubSubChannels(ctx, pattern)
}

func (cli *RedisClient) PubSubNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	return cli.client.PubSubNumSub(ctx, channels...)
}

func (cli *RedisClient) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	return cli.client.PubSubNumPat(ctx)
}

func (cli *RedisClient) PubSubShardChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return cli.client.PubSubShardChannels(ctx, pattern)
}

func (cli *RedisClient) PubSubShardNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	return cli.client.PubSubShardNumSub(ctx, channels...)
}

func (cli *RedisClient) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	return cli.client.ClusterSlots(ctx)
}

func (cli *RedisClient) ClusterNodes(ctx context.Context) *redis.StringCmd {
	return cli.client.ClusterNodes(ctx)
}

func (cli *RedisClient) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	return cli.client.ClusterMeet(ctx, host, port)
}

func (cli *RedisClient) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	return cli.client.ClusterForget(ctx, nodeID)
}

func (cli *RedisClient) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	return cli.client.ClusterReplicate(ctx, nodeID)
}

func (cli *RedisClient) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	return cli.client.ClusterResetSoft(ctx)
}

func (cli *RedisClient) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	return cli.client.ClusterResetHard(ctx)
}

func (cli *RedisClient) ClusterInfo(ctx context.Context) *redis.StringCmd {
	return cli.client.ClusterInfo(ctx)
}

func (cli *RedisClient) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	return cli.client.ClusterKeySlot(ctx, key)
}

func (cli *RedisClient) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd {
	return cli.client.ClusterGetKeysInSlot(ctx, slot, count)
}

func (cli *RedisClient) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	return cli.client.ClusterCountFailureReports(ctx, nodeID)
}

func (cli *RedisClient) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	return cli.client.ClusterCountKeysInSlot(ctx, slot)
}

func (cli *RedisClient) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	return cli.client.ClusterDelSlots(ctx, slots...)
}

func (cli *RedisClient) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	return cli.client.ClusterDelSlotsRange(ctx, min, max)
}

func (cli *RedisClient) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	return cli.client.ClusterSaveConfig(ctx)
}

func (cli *RedisClient) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	return cli.client.ClusterSlaves(ctx, nodeID)
}

func (cli *RedisClient) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	return cli.client.ClusterFailover(ctx)
}

func (cli *RedisClient) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	return cli.client.ClusterAddSlots(ctx, slots...)
}

func (cli *RedisClient) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	return cli.client.ClusterAddSlotsRange(ctx, min, max)
}

func (cli *RedisClient) GeoAdd(ctx context.Context, key string, geoLocation ...*redis.GeoLocation) *redis.IntCmd {
	return cli.client.GeoAdd(ctx, key, geoLocation...)
}

func (cli *RedisClient) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	return cli.client.GeoPos(ctx, key, members...)
}

func (cli *RedisClient) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return cli.client.GeoRadius(ctx, key, longitude, latitude, query)
}

func (cli *RedisClient) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *redis.GeoRadiusQuery) *redis.IntCmd {
	return cli.client.GeoRadiusStore(ctx, key, longitude, latitude, query)
}

func (cli *RedisClient) GeoRadiusByMember(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.GeoLocationCmd {
	return cli.client.GeoRadiusByMember(ctx, key, member, query)
}

func (cli *RedisClient) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *redis.GeoRadiusQuery) *redis.IntCmd {
	return cli.client.GeoRadiusByMemberStore(ctx, key, member, query)
}

func (cli *RedisClient) GeoSearch(ctx context.Context, key string, q *redis.GeoSearchQuery) *redis.StringSliceCmd {
	return cli.client.GeoSearch(ctx, key, q)
}

func (cli *RedisClient) GeoSearchLocation(ctx context.Context, key string, q *redis.GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	return cli.client.GeoSearchLocation(ctx, key, q)
}

func (cli *RedisClient) GeoSearchStore(ctx context.Context, key, store string, q *redis.GeoSearchStoreQuery) *redis.IntCmd {
	return cli.client.GeoSearchStore(ctx, key, store, q)
}

func (cli *RedisClient) GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd {
	return cli.client.GeoDist(ctx, key, member1, member2, unit)
}
func (cli *RedisClient) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	return cli.client.GeoHash(ctx, key, members...)
}

func (cli *RedisClient) AddHook(hook redis.Hook) {
	cli.client.AddHook(hook)
}

func (cli *RedisClient) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {
	return cli.client.Watch(ctx, fn, keys...)
}

func (cli *RedisClient) Do(ctx context.Context, args ...any) *redis.Cmd {
	return cli.client.Do(ctx, args...)
}

func (cli *RedisClient) Process(ctx context.Context, cmd redis.Cmder) error {
	return cli.client.Process(ctx, cmd)
}

func (cli *RedisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return cli.client.Subscribe(ctx, channels...)
}

func (cli *RedisClient) PSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return cli.client.PSubscribe(ctx, channels...)
}

func (cli *RedisClient) SSubscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return cli.client.SSubscribe(ctx, channels...)
}

func (cli *RedisClient) Close() error {
	return cli.client.Close()
}

func (cli *RedisClient) PoolStats() *redis.PoolStats {
	return cli.client.PoolStats()
}

func (cli *RedisClient) CommandList(ctx context.Context, filter *redis.FilterBy) *redis.StringSliceCmd {
	return cli.client.CommandList(ctx, filter)
}

func (cli *RedisClient) CommandGetKeys(ctx context.Context, commands ...any) *redis.StringSliceCmd {
	return cli.client.CommandGetKeys(ctx, commands...)
}

func (cli *RedisClient) CommandGetKeysAndFlags(ctx context.Context, commands ...any) *redis.KeyFlagsCmd {
	return cli.client.CommandGetKeysAndFlags(ctx, commands...)
}

func (cli *RedisClient) ClientInfo(ctx context.Context) *redis.ClientInfoCmd {
	return cli.client.ClientInfo(ctx)
}

func (cli *RedisClient) ModuleLoadex(ctx context.Context, conf *redis.ModuleLoadexConfig) *redis.StringCmd {
	return cli.client.ModuleLoadex(ctx, conf)
}

func (cli *RedisClient) ACLDryRun(ctx context.Context, username string, command ...any) *redis.StringCmd {
	return cli.client.ACLDryRun(ctx, username, command...)
}

func (cli *RedisClient) ACLLog(ctx context.Context, count int64) *redis.ACLLogCmd {
	return cli.client.ACLLog(ctx, count)
}

func (cli *RedisClient) ACLLogReset(ctx context.Context) *redis.StatusCmd {
	return cli.client.ACLLogReset(ctx)
}

func (cli *RedisClient) ACLSetUser(ctx context.Context, username string, rules ...string) *redis.StatusCmd {
	return cli.client.ACLSetUser(ctx, username, rules...)
}

func (cli *RedisClient) ACLDelUser(ctx context.Context, username string) *redis.IntCmd {
	return cli.client.ACLDelUser(ctx, username)
}

func (cli *RedisClient) ACLList(ctx context.Context) *redis.StringSliceCmd {
	return cli.client.ACLList(ctx)
}

func (cli *RedisClient) ACLCat(ctx context.Context) *redis.StringSliceCmd {
	return cli.client.ACLCat(ctx)
}

func (cli *RedisClient) ACLCatArgs(ctx context.Context, options *redis.ACLCatArgs) *redis.StringSliceCmd {
	return cli.client.ACLCatArgs(ctx, options)
}

func (cli *RedisClient) BitPosSpan(ctx context.Context, key string, bit int8, start int64, end int64, span string) *redis.IntCmd {
	return cli.client.BitPosSpan(ctx, key, bit, start, end, span)
}

func (cli *RedisClient) BitFieldRO(ctx context.Context, key string, values ...any) *redis.IntSliceCmd {
	return cli.client.BitFieldRO(ctx, key, values...)
}

func (cli *RedisClient) ClusterMyShardID(ctx context.Context) *redis.StringCmd {
	return cli.client.ClusterMyShardID(ctx)
}

func (cli *RedisClient) ClusterMyID(ctx context.Context) *redis.StringCmd {
	return cli.client.ClusterMyID(ctx)
}

func (cli *RedisClient) ClusterShards(ctx context.Context) *redis.ClusterShardsCmd {
	return cli.client.ClusterShards(ctx)
}

func (cli *RedisClient) ClusterLinks(ctx context.Context) *redis.ClusterLinksCmd {
	return cli.client.ClusterLinks(ctx)
}

func (cli *RedisClient) ExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	return cli.client.ExpireTime(ctx, key)
}

func (cli *RedisClient) ObjectFreq(ctx context.Context, key string) *redis.IntCmd {
	return cli.client.ObjectFreq(ctx, key)
}

func (cli *RedisClient) PExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	return cli.client.PExpireTime(ctx, key)
}

func (cli *RedisClient) SortRO(ctx context.Context, key string, sort *redis.Sort) *redis.StringSliceCmd {
	return cli.client.SortRO(ctx, key, sort)
}

func (cli *RedisClient) HGetDel(ctx context.Context, key string, fields ...string) *redis.StringSliceCmd {
	return cli.client.HGetDel(ctx, key, fields...)
}

func (cli *RedisClient) HGetEX(ctx context.Context, key string, fields ...string) *redis.StringSliceCmd {
	return cli.client.HGetEX(ctx, key, fields...)
}

func (cli *RedisClient) HGetEXWithArgs(ctx context.Context, key string, options *redis.HGetEXOptions, fields ...string) *redis.StringSliceCmd {
	return cli.client.HGetEXWithArgs(ctx, key, options, fields...)
}

func (cli *RedisClient) HSetEX(ctx context.Context, key string, fieldsAndValues ...string) *redis.IntCmd {
	return cli.client.HSetEX(ctx, key, fieldsAndValues...)
}

func (cli *RedisClient) HSetEXWithArgs(ctx context.Context, key string, options *redis.HSetEXOptions, fieldsAndValues ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HSetEXWithArgs(ctx, fullKey, options, fieldsAndValues...)
}

func (cli *RedisClient) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HScanNoValues(ctx, fullKey, cursor, match, count)
}

func (cli *RedisClient) HStrLen(ctx context.Context, key string, field string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HStrLen(ctx, fullKey, field)
}

func (cli *RedisClient) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HExpire(ctx, fullKey, expiration, fields...)
}

func (cli *RedisClient) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HExpireWithArgs(ctx, fullKey, expiration, expirationArgs, fields...)
}

func (cli *RedisClient) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HPExpire(ctx, fullKey, expiration, fields...)
}

func (cli *RedisClient) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HPExpireWithArgs(ctx, fullKey, expiration, expirationArgs, fields...)
}

func (cli *RedisClient) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HExpireAt(ctx, fullKey, tm, fields...)
}

func (cli *RedisClient) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HExpireAtWithArgs(ctx, fullKey, tm, expirationArgs, fields...)
}

func (cli *RedisClient) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HPExpireAt(ctx, fullKey, tm, fields...)
}

func (cli *RedisClient) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HPExpireAtWithArgs(ctx, fullKey, tm, expirationArgs, fields...)
}

func (cli *RedisClient) HPersist(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HPersist(ctx, fullKey, fields...)
}

func (cli *RedisClient) HExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HExpireTime(ctx, fullKey, fields...)
}

func (cli *RedisClient) HPExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HPExpireTime(ctx, fullKey, fields...)
}

func (cli *RedisClient) HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HTTL(ctx, fullKey, fields...)
}

func (cli *RedisClient) HPTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.HPTTL(ctx, fullKey, fields...)
}

func (cli *RedisClient) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	fullKeys := addPrefixToKeys(ctx, keys)
	return cli.client.BLMPop(ctx, timeout, direction, count, fullKeys...)
}

func (cli *RedisClient) LMPop(ctx context.Context, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	fullKeys := addPrefixToKeys(ctx, keys)
	return cli.client.LMPop(ctx, direction, count, fullKeys...)
}

func (cli *RedisClient) BFAdd(ctx context.Context, key string, element any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFAdd(ctx, fullKey, element)
}

func (cli *RedisClient) BFCard(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFCard(ctx, fullKey)
}

func (cli *RedisClient) BFExists(ctx context.Context, key string, element any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFExists(ctx, fullKey, element)
}

func (cli *RedisClient) BFInfo(ctx context.Context, key string) *redis.BFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInfo(ctx, fullKey)
}

func (cli *RedisClient) BFInfoArg(ctx context.Context, key string, option string) *redis.BFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInfoArg(ctx, fullKey, option)
}

func (cli *RedisClient) BFInfoCapacity(ctx context.Context, key string) *redis.BFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInfoCapacity(ctx, fullKey)
}

func (cli *RedisClient) BFInfoSize(ctx context.Context, key string) *redis.BFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInfoSize(ctx, fullKey)
}

func (cli *RedisClient) BFInfoFilters(ctx context.Context, key string) *redis.BFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInfoFilters(ctx, fullKey)
}

func (cli *RedisClient) BFInfoItems(ctx context.Context, key string) *redis.BFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInfoItems(ctx, fullKey)
}

func (cli *RedisClient) BFInfoExpansion(ctx context.Context, key string) *redis.BFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInfoExpansion(ctx, fullKey)
}

func (cli *RedisClient) BFInsert(ctx context.Context, key string, options *redis.BFInsertOptions, elements ...any) *redis.BoolSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFInsert(ctx, fullKey, options, elements...)
}

func (cli *RedisClient) BFMAdd(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFMAdd(ctx, fullKey, elements...)
}

func (cli *RedisClient) BFMExists(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFMExists(ctx, fullKey, elements...)
}

func (cli *RedisClient) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFReserve(ctx, fullKey, errorRate, capacity)
}

func (cli *RedisClient) BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFReserveNonScaling(ctx, fullKey, errorRate, capacity)
}

func (cli *RedisClient) BFReserveWithArgs(ctx context.Context, key string, options *redis.BFReserveOptions) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFReserveWithArgs(ctx, fullKey, options)
}

func (cli *RedisClient) BFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFScanDump(ctx, fullKey, iterator)
}

func (cli *RedisClient) BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity int64, expansion int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFReserveExpansion(ctx, fullKey, errorRate, capacity, expansion)
}

func (cli *RedisClient) BFLoadChunk(ctx context.Context, key string, iterator int64, data any) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.BFLoadChunk(ctx, fullKey, iterator, data)
}

func (cli *RedisClient) CFAdd(ctx context.Context, key string, element any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFAdd(ctx, fullKey, element)
}

func (cli *RedisClient) CFAddNX(ctx context.Context, key string, element any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFAddNX(ctx, fullKey, element)
}

func (cli *RedisClient) CFCount(ctx context.Context, key string, element any) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFCount(ctx, fullKey, element)
}

func (cli *RedisClient) CFDel(ctx context.Context, key string, element any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFDel(ctx, fullKey, element)
}

func (cli *RedisClient) CFExists(ctx context.Context, key string, element any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFExists(ctx, fullKey, element)
}

func (cli *RedisClient) CFInfo(ctx context.Context, key string) *redis.CFInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFInfo(ctx, fullKey)
}

func (cli *RedisClient) CFInsert(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...any) *redis.BoolSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFInsert(ctx, fullKey, options, elements...)
}

func (cli *RedisClient) CFInsertNX(ctx context.Context, key string, options *redis.CFInsertOptions, elements ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFInsertNX(ctx, fullKey, options, elements...)
}

func (cli *RedisClient) CFMExists(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFMExists(ctx, fullKey, elements...)
}

func (cli *RedisClient) CFReserve(ctx context.Context, key string, capacity int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFReserve(ctx, fullKey, capacity)
}

func (cli *RedisClient) CFReserveWithArgs(ctx context.Context, key string, options *redis.CFReserveOptions) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFReserveWithArgs(ctx, fullKey, options)
}

func (cli *RedisClient) CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFReserveExpansion(ctx, fullKey, capacity, expansion)
}

func (cli *RedisClient) CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFReserveBucketSize(ctx, fullKey, capacity, bucketsize)
}

func (cli *RedisClient) CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFReserveMaxIterations(ctx, fullKey, capacity, maxiterations)
}

func (cli *RedisClient) CFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFScanDump(ctx, fullKey, iterator)
}

func (cli *RedisClient) CFLoadChunk(ctx context.Context, key string, iterator int64, data any) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CFLoadChunk(ctx, fullKey, iterator, data)
}

func (cli *RedisClient) CMSIncrBy(ctx context.Context, key string, elements ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CMSIncrBy(ctx, fullKey, elements...)
}

func (cli *RedisClient) CMSInfo(ctx context.Context, key string) *redis.CMSInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CMSInfo(ctx, fullKey)
}

func (cli *RedisClient) CMSInitByDim(ctx context.Context, key string, width int64, height int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CMSInitByDim(ctx, fullKey, width, height)
}

func (cli *RedisClient) CMSInitByProb(ctx context.Context, key string, errorRate float64, probability float64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CMSInitByProb(ctx, fullKey, errorRate, probability)
}

func (cli *RedisClient) CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullDestKey := cachePrefix + destKey

	fullSourceKeys := make([]string, 0, len(sourceKeys))
	for _, k := range sourceKeys {
		fullSourceKeys = append(fullSourceKeys, cachePrefix+k)
	}

	return cli.client.CMSMerge(ctx, fullDestKey, fullSourceKeys...)
}

func (cli *RedisClient) CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullDestKey := cachePrefix + destKey

	fullSourceKeys := make(map[string]int64, len(sourceKeys))
	for k, v := range sourceKeys {
		fullSourceKeys[cachePrefix+k] = v
	}

	return cli.client.CMSMergeWithWeight(ctx, fullDestKey, fullSourceKeys)
}

func (cli *RedisClient) CMSQuery(ctx context.Context, key string, elements ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.CMSQuery(ctx, fullKey, elements...)
}

func (cli *RedisClient) TopKAdd(ctx context.Context, key string, elements ...any) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKAdd(ctx, fullKey, elements...)
}

func (cli *RedisClient) TopKCount(ctx context.Context, key string, elements ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKCount(ctx, fullKey, elements...)
}

func (cli *RedisClient) TopKIncrBy(ctx context.Context, key string, elements ...any) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKIncrBy(ctx, fullKey, elements...)
}

func (cli *RedisClient) TopKInfo(ctx context.Context, key string) *redis.TopKInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKInfo(ctx, fullKey)
}

func (cli *RedisClient) TopKList(ctx context.Context, key string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKList(ctx, fullKey)
}

func (cli *RedisClient) TopKListWithCount(ctx context.Context, key string) *redis.MapStringIntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKListWithCount(ctx, fullKey)
}

func (cli *RedisClient) TopKQuery(ctx context.Context, key string, elements ...any) *redis.BoolSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKQuery(ctx, fullKey, elements...)
}

func (cli *RedisClient) TopKReserve(ctx context.Context, key string, k int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKReserve(ctx, fullKey, k)
}

func (cli *RedisClient) TopKReserveWithOptions(ctx context.Context, key string, k int64, width int64, depth int64, decay float64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TopKReserveWithOptions(ctx, fullKey, k, width, depth, decay)
}

func (cli *RedisClient) TDigestAdd(ctx context.Context, key string, elements ...float64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestAdd(ctx, fullKey, elements...)
}

func (cli *RedisClient) TDigestByRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestByRank(ctx, fullKey, rank...)
}

func (cli *RedisClient) TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestByRevRank(ctx, fullKey, rank...)
}

func (cli *RedisClient) TDigestCDF(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestCDF(ctx, fullKey, elements...)
}

func (cli *RedisClient) TDigestCreate(ctx context.Context, key string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestCreate(ctx, fullKey)
}

func (cli *RedisClient) TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestCreateWithCompression(ctx, fullKey, compression)
}

func (cli *RedisClient) TDigestInfo(ctx context.Context, key string) *redis.TDigestInfoCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestInfo(ctx, fullKey)
}

func (cli *RedisClient) TDigestMax(ctx context.Context, key string) *redis.FloatCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestMax(ctx, fullKey)
}

func (cli *RedisClient) TDigestMin(ctx context.Context, key string) *redis.FloatCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestMin(ctx, fullKey)
}

func (cli *RedisClient) TDigestMerge(ctx context.Context, destKey string, options *redis.TDigestMergeOptions, sourceKeys ...string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullDestKey := cachePrefix + destKey

	fullSourceKeys := make([]string, 0, len(sourceKeys))
	for _, k := range sourceKeys {
		fullSourceKeys = append(fullSourceKeys, cachePrefix+k)
	}

	return cli.client.TDigestMerge(ctx, fullDestKey, options, fullSourceKeys...)
}

func (cli *RedisClient) TDigestQuantile(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestQuantile(ctx, fullKey, elements...)
}

func (cli *RedisClient) TDigestRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestRank(ctx, fullKey, values...)
}

func (cli *RedisClient) TDigestReset(ctx context.Context, key string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestReset(ctx, fullKey)
}

func (cli *RedisClient) TDigestRevRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestRevRank(ctx, fullKey, values...)
}

func (cli *RedisClient) TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile float64, highCutQuantile float64) *redis.FloatCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TDigestTrimmedMean(ctx, fullKey, lowCutQuantile, highCutQuantile)
}

func (cli *RedisClient) FunctionLoad(ctx context.Context, code string) *redis.StringCmd {
	return cli.client.FunctionLoad(ctx, code)
}

func (cli *RedisClient) FunctionLoadReplace(ctx context.Context, code string) *redis.StringCmd {
	return cli.client.FunctionLoadReplace(ctx, code)
}

func (cli *RedisClient) FunctionDelete(ctx context.Context, libName string) *redis.StringCmd {
	return cli.client.FunctionDelete(ctx, libName)
}

func (cli *RedisClient) FunctionFlush(ctx context.Context) *redis.StringCmd {
	return cli.client.FunctionFlush(ctx)
}

func (cli *RedisClient) FunctionKill(ctx context.Context) *redis.StringCmd {
	return cli.client.FunctionKill(ctx)
}

func (cli *RedisClient) FunctionFlushAsync(ctx context.Context) *redis.StringCmd {
	return cli.client.FunctionFlushAsync(ctx)
}

func (cli *RedisClient) FunctionList(ctx context.Context, q redis.FunctionListQuery) *redis.FunctionListCmd {
	return cli.client.FunctionList(ctx, q)
}

func (cli *RedisClient) FunctionDump(ctx context.Context) *redis.StringCmd {
	return cli.client.FunctionDump(ctx)
}

func (cli *RedisClient) FunctionRestore(ctx context.Context, libDump string) *redis.StringCmd {
	return cli.client.FunctionRestore(ctx, libDump)
}

func (cli *RedisClient) FunctionStats(ctx context.Context) *redis.FunctionStatsCmd {
	return cli.client.FunctionStats(ctx)
}

func (cli *RedisClient) FCall(ctx context.Context, function string, keys []string, args ...any) *redis.Cmd {
	return cli.client.FCall(ctx, function, keys, args...)
}

func (cli *RedisClient) FCallRo(ctx context.Context, function string, keys []string, args ...any) *redis.Cmd {
	return cli.client.FCallRo(ctx, function, keys, args...)
}

func (cli *RedisClient) FCallRO(ctx context.Context, function string, keys []string, args ...any) *redis.Cmd {
	return cli.client.FCallRo(ctx, function, keys, args...)
}

func (cli *RedisClient) FT_List(ctx context.Context) *redis.StringSliceCmd {
	return cli.client.FT_List(ctx)
}

func (cli *RedisClient) FTAggregate(ctx context.Context, index string, query string) *redis.MapStringInterfaceCmd {
	return cli.client.FTAggregate(ctx, index, query)
}

func (cli *RedisClient) FTAggregateWithArgs(ctx context.Context, index string, query string, options *redis.FTAggregateOptions) *redis.AggregateCmd {
	return cli.client.FTAggregateWithArgs(ctx, index, query, options)
}

func (cli *RedisClient) FTAliasAdd(ctx context.Context, index string, alias string) *redis.StatusCmd {
	return cli.client.FTAliasAdd(ctx, index, alias)
}

func (cli *RedisClient) FTAliasDel(ctx context.Context, alias string) *redis.StatusCmd {
	return cli.client.FTAliasDel(ctx, alias)
}

func (cli *RedisClient) FTAliasUpdate(ctx context.Context, index string, alias string) *redis.StatusCmd {
	return cli.client.FTAliasUpdate(ctx, index, alias)
}

func (cli *RedisClient) FTAlter(ctx context.Context, index string, skipInitialScan bool, definition []any) *redis.StatusCmd {
	return cli.client.FTAlter(ctx, index, skipInitialScan, definition)
}

func (cli *RedisClient) FTConfigGet(ctx context.Context, option string) *redis.MapMapStringInterfaceCmd {
	return cli.client.FTConfigGet(ctx, option)
}

func (cli *RedisClient) FTConfigSet(ctx context.Context, option string, value any) *redis.StatusCmd {
	return cli.client.FTConfigSet(ctx, option, value)
}

func (cli *RedisClient) FTCreate(ctx context.Context, index string, options *redis.FTCreateOptions, schema ...*redis.FieldSchema) *redis.StatusCmd {
	return cli.client.FTCreate(ctx, index, options, schema...)
}

func (cli *RedisClient) FTCursorDel(ctx context.Context, index string, cursorId int) *redis.StatusCmd {
	return cli.client.FTCursorDel(ctx, index, cursorId)
}

func (cli *RedisClient) FTCursorRead(ctx context.Context, index string, cursorId int, count int) *redis.MapStringInterfaceCmd {
	return cli.client.FTCursorRead(ctx, index, cursorId, count)
}

func (cli *RedisClient) FTDictAdd(ctx context.Context, dict string, term ...any) *redis.IntCmd {
	return cli.client.FTDictAdd(ctx, dict, term...)
}

func (cli *RedisClient) FTDictDel(ctx context.Context, dict string, term ...any) *redis.IntCmd {
	return cli.client.FTDictDel(ctx, dict, term...)
}

func (cli *RedisClient) FTDictDump(ctx context.Context, dict string) *redis.StringSliceCmd {
	return cli.client.FTDictDump(ctx, dict)
}

func (cli *RedisClient) FTDropIndex(ctx context.Context, index string) *redis.StatusCmd {
	return cli.client.FTDropIndex(ctx, index)
}

func (cli *RedisClient) FTDropIndexWithArgs(ctx context.Context, index string, options *redis.FTDropIndexOptions) *redis.StatusCmd {
	return cli.client.FTDropIndexWithArgs(ctx, index, options)
}

func (cli *RedisClient) FTExplain(ctx context.Context, index string, query string) *redis.StringCmd {
	return cli.client.FTExplain(ctx, index, query)
}

func (cli *RedisClient) FTExplainWithArgs(ctx context.Context, index string, query string, options *redis.FTExplainOptions) *redis.StringCmd {
	return cli.client.FTExplainWithArgs(ctx, index, query, options)
}

func (cli *RedisClient) FTInfo(ctx context.Context, index string) *redis.FTInfoCmd {
	return cli.client.FTInfo(ctx, index)
}

func (cli *RedisClient) FTSpellCheck(ctx context.Context, index string, query string) *redis.FTSpellCheckCmd {
	return cli.client.FTSpellCheck(ctx, index, query)
}

func (cli *RedisClient) FTSpellCheckWithArgs(ctx context.Context, index string, query string, options *redis.FTSpellCheckOptions) *redis.FTSpellCheckCmd {
	return cli.client.FTSpellCheckWithArgs(ctx, index, query, options)
}

func (cli *RedisClient) FTSearch(ctx context.Context, index string, query string) *redis.FTSearchCmd {
	return cli.client.FTSearch(ctx, index, query)
}

func (cli *RedisClient) FTSearchWithArgs(ctx context.Context, index string, query string, options *redis.FTSearchOptions) *redis.FTSearchCmd {
	return cli.client.FTSearchWithArgs(ctx, index, query, options)
}

func (cli *RedisClient) FTSynDump(ctx context.Context, index string) *redis.FTSynDumpCmd {
	return cli.client.FTSynDump(ctx, index)
}

func (cli *RedisClient) FTSynUpdate(ctx context.Context, index string, synGroupId any, terms []any) *redis.StatusCmd {
	return cli.client.FTSynUpdate(ctx, index, synGroupId, terms)
}

func (cli *RedisClient) FTSynUpdateWithArgs(ctx context.Context, index string, synGroupId any, options *redis.FTSynUpdateOptions, terms []any) *redis.StatusCmd {
	return cli.client.FTSynUpdateWithArgs(ctx, index, synGroupId, options, terms)
}

func (cli *RedisClient) FTTagVals(ctx context.Context, index string, field string) *redis.StringSliceCmd {
	return cli.client.FTTagVals(ctx, index, field)
}

func (cli *RedisClient) SInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		fullKeys = append(fullKeys, cachePrefix+k)
	}
	return cli.client.SInterCard(ctx, limit, fullKeys...)
}

func (cli *RedisClient) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		fullKeys = append(fullKeys, cachePrefix+k)
	}
	return cli.client.BZMPop(ctx, timeout, order, count, fullKeys...)
}

func (cli *RedisClient) ZAddLT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZAddLT(ctx, fullKey, members...)
}

func (cli *RedisClient) ZAddGT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZAddGT(ctx, fullKey, members...)
}

func (cli *RedisClient) ZInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		fullKeys = append(fullKeys, cachePrefix+k)
	}
	return cli.client.ZInterCard(ctx, limit, fullKeys...)
}

func (cli *RedisClient) ZMPop(ctx context.Context, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		fullKeys = append(fullKeys, cachePrefix+k)
	}
	return cli.client.ZMPop(ctx, order, count, fullKeys...)
}

func (cli *RedisClient) ZRankWithScore(ctx context.Context, key string, member string) *redis.RankWithScoreCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRankWithScore(ctx, fullKey, member)
}

func (cli *RedisClient) ZRevRankWithScore(ctx context.Context, key string, member string) *redis.RankWithScoreCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.ZRevRankWithScore(ctx, fullKey, member)
}

func (cli *RedisClient) LCS(ctx context.Context, q *redis.LCSQuery) *redis.LCSCmd {
	return cli.client.LCS(ctx, q)
}

func (cli *RedisClient) TSAdd(ctx context.Context, key string, timestamp any, value float64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSAdd(ctx, fullKey, timestamp, value)
}

func (cli *RedisClient) TSAddWithArgs(ctx context.Context, key string, timestamp any, value float64, options *redis.TSOptions) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSAddWithArgs(ctx, fullKey, timestamp, value, options)
}

func (cli *RedisClient) TSCreate(ctx context.Context, key string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSCreate(ctx, fullKey)
}

func (cli *RedisClient) TSCreateWithArgs(ctx context.Context, key string, options *redis.TSOptions) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSCreateWithArgs(ctx, fullKey, options)
}

func (cli *RedisClient) TSAlter(ctx context.Context, key string, options *redis.TSAlterOptions) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSAlter(ctx, fullKey, options)
}

func (cli *RedisClient) TSCreateRule(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullSourceKey := cachePrefix + sourceKey
	fullDestKey := cachePrefix + destKey
	return cli.client.TSCreateRule(ctx, fullSourceKey, fullDestKey, aggregator, bucketDuration)
}

func (cli *RedisClient) TSCreateRuleWithArgs(ctx context.Context, sourceKey string, destKey string, aggregator redis.Aggregator, bucketDuration int, options *redis.TSCreateRuleOptions) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullSourceKey := cachePrefix + sourceKey
	fullDestKey := cachePrefix + destKey
	return cli.client.TSCreateRuleWithArgs(ctx, fullSourceKey, fullDestKey, aggregator, bucketDuration, options)
}

func (cli *RedisClient) TSIncrBy(ctx context.Context, key string, timestamp float64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSIncrBy(ctx, fullKey, timestamp)
}

func (cli *RedisClient) TSIncrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSIncrByWithArgs(ctx, fullKey, timestamp, options)
}

func (cli *RedisClient) TSDecrBy(ctx context.Context, key string, timestamp float64) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSDecrBy(ctx, fullKey, timestamp)
}

func (cli *RedisClient) TSDecrByWithArgs(ctx context.Context, key string, timestamp float64, options *redis.TSIncrDecrOptions) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSDecrByWithArgs(ctx, fullKey, timestamp, options)
}

func (cli *RedisClient) TSDel(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSDel(ctx, fullKey, fromTimestamp, toTimestamp)
}

func (cli *RedisClient) TSDeleteRule(ctx context.Context, sourceKey string, destKey string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	srcFullKey := cachePrefix + sourceKey
	destFullKey := cachePrefix + destKey
	return cli.client.TSDeleteRule(ctx, srcFullKey, destFullKey)
}

func (cli *RedisClient) TSGet(ctx context.Context, key string) *redis.TSTimestampValueCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSGet(ctx, fullKey)
}

func (cli *RedisClient) TSGetWithArgs(ctx context.Context, key string, options *redis.TSGetOptions) *redis.TSTimestampValueCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSGetWithArgs(ctx, fullKey, options)
}

func (cli *RedisClient) TSInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSInfo(ctx, fullKey)
}

func (cli *RedisClient) TSInfoWithArgs(ctx context.Context, key string, options *redis.TSInfoOptions) *redis.MapStringInterfaceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSInfoWithArgs(ctx, fullKey, options)
}

func (cli *RedisClient) TSMAdd(ctx context.Context, ktvSlices [][]any) *redis.IntSliceCmd {
	return cli.client.TSMAdd(ctx, ktvSlices)
}

func (cli *RedisClient) TSQueryIndex(ctx context.Context, filterExpr []string) *redis.StringSliceCmd {
	return cli.client.TSQueryIndex(ctx, filterExpr)
}

func (cli *RedisClient) TSRevRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSRevRange(ctx, fullKey, fromTimestamp, toTimestamp)
}

func (cli *RedisClient) TSRevRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRevRangeOptions) *redis.TSTimestampValueSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSRevRangeWithArgs(ctx, fullKey, fromTimestamp, toTimestamp, options)
}

func (cli *RedisClient) TSRange(ctx context.Context, key string, fromTimestamp int, toTimestamp int) *redis.TSTimestampValueSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSRange(ctx, fullKey, fromTimestamp, toTimestamp)
}

func (cli *RedisClient) TSRangeWithArgs(ctx context.Context, key string, fromTimestamp int, toTimestamp int, options *redis.TSRangeOptions) *redis.TSTimestampValueSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.TSRangeWithArgs(ctx, fullKey, fromTimestamp, toTimestamp, options)
}

func (cli *RedisClient) TSMRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	return cli.client.TSMRange(ctx, fromTimestamp, toTimestamp, filterExpr)
}

func (cli *RedisClient) TSMRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRangeOptions) *redis.MapStringSliceInterfaceCmd {
	return cli.client.TSMRangeWithArgs(ctx, fromTimestamp, toTimestamp, filterExpr, options)
}

func (cli *RedisClient) TSMRevRange(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string) *redis.MapStringSliceInterfaceCmd {
	return cli.client.TSMRevRange(ctx, fromTimestamp, toTimestamp, filterExpr)
}

func (cli *RedisClient) TSMRevRangeWithArgs(ctx context.Context, fromTimestamp int, toTimestamp int, filterExpr []string, options *redis.TSMRevRangeOptions) *redis.MapStringSliceInterfaceCmd {
	return cli.client.TSMRevRangeWithArgs(ctx, fromTimestamp, toTimestamp, filterExpr, options)
}

func (cli *RedisClient) TSMGet(ctx context.Context, filters []string) *redis.MapStringSliceInterfaceCmd {
	return cli.client.TSMGet(ctx, filters)
}

func (cli *RedisClient) TSMGetWithArgs(ctx context.Context, filters []string, options *redis.TSMGetOptions) *redis.MapStringSliceInterfaceCmd {
	return cli.client.TSMGetWithArgs(ctx, filters, options)
}

func (cli *RedisClient) JSONArrAppend(ctx context.Context, key string, path string, values ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrAppend(ctx, fullKey, path, values...)
}

func (cli *RedisClient) JSONArrIndex(ctx context.Context, key string, path string, value ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrIndex(ctx, fullKey, path, value...)
}

func (cli *RedisClient) JSONArrIndexWithArgs(ctx context.Context, key string, path string, options *redis.JSONArrIndexArgs, value ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrIndexWithArgs(ctx, fullKey, path, options, value...)
}

func (cli *RedisClient) JSONArrInsert(ctx context.Context, key string, path string, index int64, values ...any) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrInsert(ctx, fullKey, path, index, values...)
}

func (cli *RedisClient) JSONArrLen(ctx context.Context, key string, path string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrLen(ctx, fullKey, path)
}

func (cli *RedisClient) JSONArrPop(ctx context.Context, key string, path string, index int) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrPop(ctx, fullKey, path, index)
}

func (cli *RedisClient) JSONArrTrim(ctx context.Context, key string, path string) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrTrim(ctx, fullKey, path)
}

func (cli *RedisClient) JSONArrTrimWithArgs(ctx context.Context, key string, path string, options *redis.JSONArrTrimArgs) *redis.IntSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONArrTrimWithArgs(ctx, fullKey, path, options)
}

func (cli *RedisClient) JSONClear(ctx context.Context, key string, path string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONClear(ctx, fullKey, path)
}

func (cli *RedisClient) JSONDebugMemory(ctx context.Context, key string, path string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONDebugMemory(ctx, fullKey, path)
}

func (cli *RedisClient) JSONDel(ctx context.Context, key string, path string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONDel(ctx, fullKey, path)
}

func (cli *RedisClient) JSONForget(ctx context.Context, key string, path string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONForget(ctx, fullKey, path)
}

func (cli *RedisClient) JSONGet(ctx context.Context, key string, paths ...string) *redis.JSONCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONGet(ctx, fullKey, paths...)
}

func (cli *RedisClient) JSONGetWithArgs(ctx context.Context, key string, options *redis.JSONGetArgs, paths ...string) *redis.JSONCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONGetWithArgs(ctx, fullKey, options, paths...)
}

func (cli *RedisClient) JSONMerge(ctx context.Context, key string, path string, value string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONMerge(ctx, fullKey, path, value)
}

func (cli *RedisClient) JSONMSetArgs(ctx context.Context, docs []redis.JSONSetArgs) *redis.StatusCmd {
	return cli.client.JSONMSetArgs(ctx, docs)
}

func (cli *RedisClient) JSONMSet(ctx context.Context, params ...any) *redis.StatusCmd {
	return cli.client.JSONMSet(ctx, params...)
}

func (cli *RedisClient) JSONMGet(ctx context.Context, path string, keys ...string) *redis.JSONSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKeys := make([]string, 0, len(keys))
	for _, k := range keys {
		fullKeys = append(fullKeys, cachePrefix+k)
	}
	return cli.client.JSONMGet(ctx, path, fullKeys...)
}

func (cli *RedisClient) JSONNumIncrBy(ctx context.Context, key string, path string, value float64) *redis.JSONCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONNumIncrBy(ctx, fullKey, path, value)
}

func (cli *RedisClient) JSONObjKeys(ctx context.Context, key string, path string) *redis.SliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONObjKeys(ctx, fullKey, path)
}

func (cli *RedisClient) JSONObjLen(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONObjLen(ctx, fullKey, path)
}

func (cli *RedisClient) JSONSet(ctx context.Context, key string, path string, value any) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONSet(ctx, fullKey, path, value)
}

func (cli *RedisClient) JSONSetMode(ctx context.Context, key string, path string, value any, mode string) *redis.StatusCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONSetMode(ctx, fullKey, path, value, mode)
}

func (cli *RedisClient) JSONStrAppend(ctx context.Context, key string, path string, value string) *redis.IntPointerSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONStrAppend(ctx, fullKey, path, value)
}

func (cli *RedisClient) JSONStrLen(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONStrLen(ctx, fullKey, path)
}

func (cli *RedisClient) JSONToggle(ctx context.Context, key string, path string) *redis.IntPointerSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONToggle(ctx, fullKey, path)
}

func (cli *RedisClient) JSONType(ctx context.Context, key string, path string) *redis.JSONSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.JSONType(ctx, fullKey, path)
}

func (cli *RedisClient) VAdd(ctx context.Context, key string, element string, val redis.Vector) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VAdd(ctx, fullKey, element, val)
}

func (cli *RedisClient) VAddWithArgs(ctx context.Context, key string, element string, val redis.Vector, addArgs *redis.VAddArgs) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VAddWithArgs(ctx, fullKey, element, val, addArgs)
}

func (cli *RedisClient) VCard(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VCard(ctx, fullKey)
}

func (cli *RedisClient) VDim(ctx context.Context, key string) *redis.IntCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VDim(ctx, fullKey)
}

func (cli *RedisClient) VEmb(ctx context.Context, key string, element string, raw bool) *redis.SliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VEmb(ctx, fullKey, element, raw)
}

func (cli *RedisClient) VGetAttr(ctx context.Context, key string, element string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VGetAttr(ctx, fullKey, element)
}

func (cli *RedisClient) VInfo(ctx context.Context, key string) *redis.MapStringInterfaceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VInfo(ctx, fullKey)
}

func (cli *RedisClient) VLinks(ctx context.Context, key string, element string) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VLinks(ctx, fullKey, element)
}

func (cli *RedisClient) VLinksWithScores(ctx context.Context, key string, element string) *redis.VectorScoreSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VLinksWithScores(ctx, fullKey, element)
}

func (cli *RedisClient) VRandMember(ctx context.Context, key string) *redis.StringCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VRandMember(ctx, fullKey)
}

func (cli *RedisClient) VRandMemberCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VRandMemberCount(ctx, fullKey, count)
}

func (cli *RedisClient) VRem(ctx context.Context, key string, element string) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VRem(ctx, fullKey, element)
}

func (cli *RedisClient) VSetAttr(ctx context.Context, key string, element string, attr any) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VSetAttr(ctx, fullKey, element, attr)
}

func (cli *RedisClient) VClearAttributes(ctx context.Context, key string, element string) *redis.BoolCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VClearAttributes(ctx, fullKey, element)
}

func (cli *RedisClient) VSim(ctx context.Context, key string, val redis.Vector) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VSim(ctx, fullKey, val)
}

func (cli *RedisClient) VSimWithScores(ctx context.Context, key string, val redis.Vector) *redis.VectorScoreSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VSimWithScores(ctx, fullKey, val)
}

func (cli *RedisClient) VSimWithArgs(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.StringSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VSimWithArgs(ctx, fullKey, val, args)
}

func (cli *RedisClient) VSimWithArgsWithScores(ctx context.Context, key string, val redis.Vector, args *redis.VSimArgs) *redis.VectorScoreSliceCmd {
	cachePrefix := GetCachePrefixFromContext(ctx)
	fullKey := cachePrefix + key
	return cli.client.VSimWithArgsWithScores(ctx, fullKey, val, args)
}
