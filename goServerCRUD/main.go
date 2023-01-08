package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	
}

var movies []Movie

// middleware
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	// because we will be passing the entire movie slice
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	// we are sending POST request, with all key values
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// creation of movie id
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)

}
func updateMovie(w http.ResponseWriter, r *http.Request){
	// set json content type
	w.Header().Set("Content-Type","application/json")
	// params
 	params := mux.Vars(r)

	// loop over the movies, range
	for index,item := range movies{
		if item.ID == params["id"]{
			// delete the movie with the id that we have sent
			movies = append(movies[:index],movies[index+1:]...)

			// add a new movie - the movie that we send in the body of postman 
			var movie Movie 
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			
			// returning to frontend
			json.NewEncoder(w).Encode(movie)

			return
		}
	}

}
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	// to get the request queries
	params := mux.Vars(r)
	for index,item := range movies {

		if item.ID  == params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func main(){
	r := mux.NewRouter()
	
	// dummy data
	movies = append(movies, Movie{ID:"1",Isbn:"4399989",Title: "Movie 1", Director: &Director{Firstname: "A1",Lastname: "B1"} })
	movies = append(movies, Movie{ID: "2",Isbn: "342312131",Title: "Movie 2",Director: &Director{Firstname: "A2",Lastname: "B2"}})
	movies = append(movies, Movie{ID: "3",Isbn: "1213216",Title: "Movie 3", Director: &Director{Firstname: "A3",Lastname: "B3"}})		

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}