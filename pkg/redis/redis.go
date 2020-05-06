package redis

import (
	"time"
)

// Redis is abstraction for redis library
type Redis interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}
