package main

import (
	"encoding/json"
	"fmt"
	"github/gorilla/mux" // https://pkg.go.dev/github.com/gorilla/mux
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json: "id"`
	ISBN     string    `json: "isbn"`
	title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	firstName string `json: "firstName`
	lastName  string `json: "lastName"`
}

var movies = []Movie{}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	movies = append(movies,
		Movie{
			ID:    "1",
			ISBN:  "3232",
			title: "Batman(1989)",
			Director: &Director{
				firstName: "Tim",
				lastName:  "Burton",
			},
		})
	movies = append(movies,
		Movie{
			ID:    "2",
			ISBN:  "3233",
			title: "Batman Returns",
			Director: &Director{
				firstName: "Tim",
				lastName:  "Burton",
			},
		})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("movies/{id}", getMovie).Methods("GET")
	r.handleFunc("/movies", createMovie).Methods("POST")
	r.handleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.handleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
