package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // https://pkg.go.dev/github.com/gorilla/mux
)

type Movie struct {
	ID       string   `json:"id"`       // key must have first character as upper case
	ISBN     string   `json:"isbn"`     // key must have first character as upper case
	Title    string   `json:"title"`    // key must have first character as upper case
	Director Director `json:"director"` // key must have first character as upper case
}

type Director struct {
	FirstName string `json:"firstName"` // key must have first character as upper case
	LastName  string `json:"lastName"`  // key must have first character as upper case
}

var movies = []Movie{}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = params["id"]
			movies = append(movies, updatedMovie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)

}

func main() {
	movies = append(movies,
		Movie{
			ID:    "1",
			ISBN:  "3232",
			Title: "Batman(1989)",
			Director: Director{
				FirstName: "Tim",
				LastName:  "Burton",
			},
		})
	movies = append(movies,
		Movie{
			ID:    "2",
			ISBN:  "3233",
			Title: "Batman Returns",
			Director: Director{
				FirstName: "Tim",
				LastName:  "Burton",
			},
		})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
