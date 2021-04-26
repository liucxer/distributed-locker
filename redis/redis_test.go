package redis

import (
	"fmt"
	lib "github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedisLocker_Lock(t *testing.T) {
	dialFunc := func() (c lib.Conn, err error) {
		c, err = lib.Dial(
			"tcp",
			fmt.Sprintf("%s:%d", "172.16.0.155", 6379),
			lib.DialPassword("fawKCeUNuEcK8mKG"),
			lib.DialDatabase(10),
		)
		return
	}

	pool := &lib.Pool{
		Dial:        dialFunc,
		MaxIdle:     3,
		MaxActive:   5,
		IdleTimeout: time.Duration(240 * time.Second),
		Wait:        true,
	}

	client := NewRedisLocker(pool)
	err := client.Lock("a", 10*time.Second)
	require.NoError(t, err)

	err = client.Lock("a", 10*time.Second)
	require.NoError(t, err)
}
