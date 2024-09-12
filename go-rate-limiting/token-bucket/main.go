package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Message{
		Status: "Successful",
		Body:   "Hi! You've reached the API. How may I help you?",
	})
}

func main() {
	http.Handle("/ping", rateLimiter(endpointHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
