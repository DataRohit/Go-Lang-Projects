package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func rateLimiter(next http.HandlerFunc) http.Handler {
	limiter := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&Message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later.",
			})
			return
		}
		next(w, r)
	})
}
