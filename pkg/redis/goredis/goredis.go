package goredis

import (
	"time"

	grds "github.com/go-redis/redis/v7"
	"github.com/nydan/glean/pkg/redis"
)

// Config for goredis connection
type Config struct {
	Addr        string
	Timeout     int
	ReadTimeout int
	PoolSize    int
	MinIdleConn int
}

type rds struct {
	c *grds.Client
}

// Connect connects redis
func Connect(cfg Config) redis.Redis {
	c := grds.NewClient(&grds.Options{
		Network:      "tcp",
		Addr:         cfg.Addr,
		MaxRetries:   2,
		DialTimeout:  time.Duration(cfg.Timeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConn,
	})
	return &rds{c: c}
}

// Get runs GET command
func (r *rds) Get(key string) (string, error) {
	return r.c.Get(key).Result()
}

// Set run SETEX command.
// Use expiration as 0 when no expiration needed.
func (r *rds) Set(key string, value interface{}, expiration time.Duration) error {
	return r.c.Set(key, value, expiration).Err()
}
