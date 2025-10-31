package ratelimiter

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type FixedWindowRateLimiter struct {
	mu     sync.RWMutex
	client *redis.Client
	limit  int
	window time.Duration
}

func NewFixedWindowRateLimiter(client *redis.Client, limit int, window time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		client: client,
		limit:  limit,
		window: window,
	}
}

func (l *FixedWindowRateLimiter) Allow(ctx context.Context, ip string) bool {
	key := "rate:" + ip

	fmt.Println(key)
	count, err := l.client.Incr(ctx, key).Result()
	if err != nil {
		return false
	}

	if count == 1 {
		l.client.Expire(ctx, key, l.window)
	}

	return count <= int64(l.limit)
}
