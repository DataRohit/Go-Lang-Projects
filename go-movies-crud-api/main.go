package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
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

var movies []Movie

func logRequest(r *http.Request) {
	log.Printf("Method: %s, URL: %s, RemoteAddr: %s", r.Method, r.URL.Path, r.RemoteAddr)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		log.Printf("Error encoding movies: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for _, movie := range movies {
		if movie.ID == id {
			if err := json.NewEncoder(w).Encode(movie); err != nil {
				log.Printf("Error encoding movie: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	log.Printf("Movie not found: %s", id)
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Title: "Dune: Part Two",
		Directors: []*Director{
			{FirstName: "Denis", LastName: "Villeneuve"},
		},
		Genre:  "Science Fiction, Adventure",
		Budget: 190_000_000.00,
	})
	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Title: "The Batman",
		Directors: []*Director{
			{FirstName: "Matt", LastName: "Reeves"},
		},
		Genre:  "Crime, Mystery, Thriller",
		Budget: 185_000_000.00,
	})
	movies = append(movies, Movie{
		ID:    uuid.New().String(),
		Title: "Avengers: Endgame",
		Directors: []*Director{
			{FirstName: "Anthony", LastName: "Russo"},
			{FirstName: "Joe", LastName: "Russo"},
		},
		Genre:  "Adventure, Science Fiction, Action",
		Budget: 356_000_000.00,
	})

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
