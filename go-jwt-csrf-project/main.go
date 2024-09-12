package main

import (
	"fmt"
	"log"

	"github.com/datarohit/go-jwt-csrf-project/db"
	"github.com/datarohit/go-jwt-csrf-project/server"
	"github.com/datarohit/go-jwt-csrf-project/server/middleware/myJwt"
)

const (
	host = "localhost"
	port = "9000"
)

func main() {
	if err := setup(); err != nil {
		log.Fatalf("Setup failed: %v", err)
	}

	if err := server.StartServer(host, port); err != nil {
		log.Fatalf("Failed to start server on %s:%s: %v", host, port, err)
	}
}

func setup() error {
	db.InitDB()

	if err := myJwt.InitJWT(); err != nil {
		return fmt.Errorf("JWT initialization failed: %w", err)
	}

	return nil
}
