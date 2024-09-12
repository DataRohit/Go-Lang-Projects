package main

import (
	"log"

	"github.com/datarohit/go-jwt-csrf-project/server"
)

const (
	host = "localhost"
	port = "9000"
)

func main() {

	if err := server.StartServer(host, port); err != nil {
		log.Fatalf("Failed to start server on %s:%s: %v", host, port, err)
	}
}
