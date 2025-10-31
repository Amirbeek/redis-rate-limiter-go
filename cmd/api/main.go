package main

import (
	"RateLimitor-with-go/internal/env"
	"RateLimitor-with-go/internal/ratelimiter"
	_ "RateLimitor-with-go/internal/ratelimiter"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {

	cfg := config{
		Addr: env.GetString("PORT", ":8080"),
		redisConfig: redisConfig{
			addr: env.GetString("REDIS_ADDR", ":6379"),
			db:   env.GetInt("REDIS_DB", 0),
			pw:   env.GetString("REDIS_PASSWORD", ""),
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 5),
			TimeFrame:            time.Second * 5,
		},
	}

	var rdb *redis.Client
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.redisConfig.addr,
		Password: cfg.redisConfig.pw,
		DB:       cfg.redisConfig.db,
	})

	rateLimiter := ratelimiter.NewFixedWindowRateLimiter(
		rdb,
		cfg.rateLimiter.RequestsPerTimeFrame,
		cfg.rateLimiter.TimeFrame)

	// Cache layer
	app := &application{
		config:      cfg,
		rateLimiter: rateLimiter,
	}

	log.Fatal(app.run())
}
