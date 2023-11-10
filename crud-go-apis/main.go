package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type movie struct {
	id       string `json: "id"`
	title    string `json:"title"`
	director string `json:"director"`
}

var movies []movie

func main() {
	fmt.Println("Welcome to CRUD Go API's")
	movies = append(movies, movie{id: "1", title: "Forrest Gump", director: "Robert Zemeckis"})
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("crud go apis server is started at port 8080")
	} else {
		log.Fatal(err)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {

}

func getMovie(w http.ResponseWriter, r *http.Request) {

}
func createMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request)
