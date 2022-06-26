package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github/gorilla/mux" // https://pkg.go.dev/github.com/gorilla/mux
)

type Movie struct{
	ID string `json: "id"`
	ISBN string `json: "isbn"`
	title string `json: "title"`
	Director *Director `json: "director"`
}

type Director struct{
	firstName string `json: "firstName`
	lastName string `json: "lastName"`
}

var movies := make([]Movies{})
 
func main() {
	r := mux.NewRouter() 
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("movies/{id}", getMovie).Methods("GET")
	r.handleFunc("/movies", createMovie).Methods("POST")
	r.handleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.handleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

}
