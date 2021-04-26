package redis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisLocker struct {
	pool *redis.Pool
}

func NewRedisLocker(pool *redis.Pool) *RedisLocker {
	return &RedisLocker{
		pool: pool,
	}
}

func (locker *RedisLocker) Lock(key string, expiresIn time.Duration) error {
	if locker.pool == nil {
		return errors.New("pool is nil")
	}

	conn := locker.pool.Get()
	defer func() { _ = conn.Close() }()
	result, err := conn.Do("SET", key, "ok", "EX", int(expiresIn/time.Second), "NX")
	if err != nil {
		return err
	}

	if result != nil {
		return errors.New("lock error")
	}

	return nil
}

func (locker *RedisLocker) Unlock(key string) error {
	if locker.pool == nil {
		return errors.New("pool is nil")
	}

	conn := locker.pool.Get()
	defer func() { _ = conn.Close() }()

	_, err := conn.Do("DEL", key)
	return err
}
