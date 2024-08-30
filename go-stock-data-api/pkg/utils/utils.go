package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func LogRequest(r *http.Request) {
	log.Printf("Method: %s, URL: %s, RemoteAddr: %s", r.Method, r.URL.Path, r.RemoteAddr)
}

func LogError(r *http.Request, msg string, err error) {
	log.Printf("%s: %v | Method: %s | URL: %s | RemoteAddr: %s", msg, err, r.Method, r.URL.Path, r.RemoteAddr)
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to write JSON response: %v", err)
	}
}

func ParseBody(r *http.Request, x interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		return err
	}

	err = json.Unmarshal(body, x)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return err
	}

	return nil
}
