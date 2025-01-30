package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"	
)

func TestGetMovies(t *testing.T) {
	// Initialize the movies slice with dummy data
	movies = []Movie{
		{ID: "1", Isbn: "12345", Title: "Test Movie 1", Director: &Director{Firstname: "Robert", Lastname: "Griesemer"}},
		{ID: "2", Isbn: "67890", Title: "Test Movie 2", Director: &Director{Firstname: "Rob", Lastname: "Pike"}},
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	router := mux.NewRouter()
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Decode the response body
	var responseMovies []Movie
	err = json.NewDecoder(rr.Body).Decode(&responseMovies)
	if err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	// Ensure the response contains the dummy data
	if len(responseMovies) != len(movies) {
		t.Errorf("Expected %d movies, got %d", len(movies), len(responseMovies))
	}
}


// Test CreateMovie handler
func TestCreateMovie(t *testing.T) {
	newMovie := Movie{
		Isbn:     "123456",
		Title:    "Test Movie",
		Director: &Director{Firstname: "Robert", Lastname: "Griesemer"},
	}

	body, err := json.Marshal(newMovie)
	if err != nil {
		t.Fatalf("Failed to marshal new movie: %v", err)
	}

	req, err := http.NewRequest("POST", "/movies", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	var createdMovie Movie
	err = json.NewDecoder(rr.Body).Decode(&createdMovie)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if createdMovie.Title != newMovie.Title || createdMovie.Director.Firstname != newMovie.Director.Firstname {
		t.Errorf("Created movie does not match input: got %+v, want %+v", createdMovie, newMovie)
	}
}

// Test DeleteMovie handler
func TestDeleteMovie(t *testing.T) {
	movies = append(movies, Movie{ID: "999", Isbn: "test", Title: "Delete Test Movie", Director: &Director{Firstname: "Test", Lastname: "Delete"}})

	req, err := http.NewRequest("DELETE", "/movie/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check if movie is deleted
	for _, movie := range movies {
		if movie.ID == "999" {
			t.Errorf("Movie with ID 999 was not deleted")
		}
	}
}
