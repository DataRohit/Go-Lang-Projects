package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func perClientRateLimiter(next http.HandlerFunc) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	cleanup := func() {
		mu.Lock()
		defer mu.Unlock()
		for ip, c := range clients {
			if time.Since(c.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
	}

	go func() {
		for range time.Tick(time.Minute) {
			cleanup()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &client{limiter: rate.NewLimiter(2, 4)}
		}
		clients[ip].lastSeen = time.Now()
		limiter := clients[ip].limiter
		mu.Unlock()

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

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Message{
		Status: "Successful",
		Body:   "Hi! You've reached the API. How may I help you?",
	})
}

func main() {
	http.Handle("/ping", perClientRateLimiter(endpointHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
