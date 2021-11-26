package router

import (
	"github.com/Aniket762/namaste-go/dbapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies",controller.GetAllMoviesController).Methods("GET")
	r.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}",controller.MarkedAsWatch).Methods("PUT")
	r.HandleFunc("/api/movie/{id}",controller.DeleteOneMovieController).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovie",controller.DeleteAllMovieController).Methods("DELETE")
	return r
}