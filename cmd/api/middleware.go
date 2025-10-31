package main

import (
	"fmt"
	"net/http"
)

func (app *application) RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allow := app.rateLimiter.Allow(r.Context(), r.RemoteAddr); !allow {
			fmt.Println("Rate limit exceeded")
			return
		}
		next.ServeHTTP(w, r)
	})
}
