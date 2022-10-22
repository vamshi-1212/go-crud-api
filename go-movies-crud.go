package main

import (
	"encodind/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:id`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder().Encoder(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder().Encoder(movies)
}
func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.body).Decode(&movie)
	movie.ID = stringconv(rand.Intn(1000000000))
	movie = append(movies, movie)
	json.Encoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("content-type", "application/json")
	//params
	params := mux.Vars(r)
	//looping
	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.body).Decoder(&movie)
			movie.ID = strngconv(rand.Intn(1000000000))
			movie = append(movies, movie)
			json.NewEncoder(w).Encoder(movie)
		}

	}

}

//delete a movie that we hvae sent
//add new movie with id that we send to the postman

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: " 1 ", Isbn: "438221", Title: "Movie one", Director: &Director{Firstname: "vamshi", Lastname: "krishna"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438253", Title: "Movie two", Director: &Director{Firstname: "Rohit", Lastname: "sharma"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandleFunc("/movies/{id)", deleteMovie).Method("DELETE")

	fmt.Printf("starating servr at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
