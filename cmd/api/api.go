package main

import (
	"RateLimitor-with-go/internal/ratelimiter"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config      config
	rateLimiter ratelimiter.Limiter
}

type redisConfig struct {
	addr string
	pw   string
	db   int
}
type config struct {
	Addr        string
	redisConfig redisConfig
	rateLimiter ratelimiter.Config
}

func (app *application) mount(addr string) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Use(app.RateLimiterMiddleware)

	r.Get("/", app.Info)
	r.Get("/health", app.health)
	return r
}

func (app *application) run() error {
	r := app.mount(app.config.Addr)

	srv := http.Server{
		Addr:         app.config.Addr,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Listening on " + app.config.Addr)
	return srv.ListenAndServe()
}
