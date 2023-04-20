package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Latname   string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func main() {
	movies = append(movies, Movie{ID: "1", Isbn: "487391", Title: "First Movie", Director: &Director{Firstname: "Matt", Latname: "Mark"}})
	movies = append(movies, Movie{ID: "2", Isbn: "463884", Title: "Second Movie", Director: &Director{Firstname: "Luke", Latname: "John"}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
