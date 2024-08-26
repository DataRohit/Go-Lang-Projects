package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID        string      `json:"id"`
	Title     string      `json:"title"`
	Directors []*Director `json:"directors"`
	Genre     string      `json:"genre"`
	Budget    float64     `json:"budget"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	router := mux.NewRouter()

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
