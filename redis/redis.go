package redis

import (
	"errors"
	"time"
)

type RedisConner interface {
	Do(cmdName string, args ...interface{}) (reply interface{}, err error)
}

type RedisLocker struct {
	conn RedisConner
}

func NewRedisLocker(conn RedisConner) *RedisLocker {
	return &RedisLocker{
		conn: conn,
	}
}

func (locker *RedisLocker) Lock(key string, expiresIn time.Duration) error {
	result, err := locker.conn.Do("SET", key, "ok", "EX", int(expiresIn/time.Second), "NX")
	if err != nil {
		return err
	}

	if result != nil {
		return errors.New("lock error")
	}

	return nil
}

func (locker *RedisLocker) Unlock(key string) error {
	_, err := locker.conn.Do("DEL", key)
	return err
}
