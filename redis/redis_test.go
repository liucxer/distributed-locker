package redis

import (
	"fmt"
	lib "github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRedisLocker_Lock(t *testing.T) {
	redisClient, err := lib.Dial(
		"tcp",
		fmt.Sprintf("%s:%d", "172.16.0.155", 6379),
		lib.DialPassword("fawKCeUNuEcK8mKG"),
		lib.DialDatabase(0),
	)

	client := NewRedisLocker(redisClient)
	err = client.Lock("a", 10*time.Second)
	require.NoError(t, err)

	err = client.Lock("a", 10*time.Second)
	require.NoError(t, err)
}
