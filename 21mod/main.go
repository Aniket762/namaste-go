package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	fmt.Println("Getting into mod 🙏")
	r := mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("GET")

	// running server
	//http.ListenAndServe(":4000",r)

	// for logging error using log
	log.Fatal(http.ListenAndServe(":8081",r))
}


func serveHome(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("<h1>Exploring mod with mux</h1>"))
}