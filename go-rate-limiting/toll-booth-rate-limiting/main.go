package main

import (
	"encoding/json"
	"log"
	"net/http"

	tollbooth "github.com/didip/tollbooth/v7"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(&Message{
		Status: "Successful",
		Body:   "Hi! You've reached the API. How may I help you?",
	})
}

func main() {
	jsonMessage, _ := json.Marshal(&Message{
		Status: "Request Failed",
		Body:   "The API is at capacity, try again later.",
	})

	limiter := tollbooth.NewLimiter(1, nil)
	limiter.SetMessageContentType("application/json")
	limiter.SetMessage(string(jsonMessage))

	http.Handle("/ping", tollbooth.LimitFuncHandler(limiter, endpointHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
