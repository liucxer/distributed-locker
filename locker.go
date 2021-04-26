package distributed_locker

import "time"

type DistributedLocker interface {
	Lock(key string, expiresIn time.Duration) error
	Unlock(key string) error
}
