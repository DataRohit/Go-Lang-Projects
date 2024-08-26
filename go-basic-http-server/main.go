package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello World!"); err != nil {
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	if _, err := fmt.Fprintf(w, "POST request successful\n"); err != nil {
		return
	}

	name := r.FormValue("name")
	address := r.FormValue("address")

	if _, err := fmt.Fprintf(w, "Name = %s\n", name); err != nil {
		return
	}
	if _, err := fmt.Fprintf(w, "Address = %s\n", address); err != nil {
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server is listening at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
