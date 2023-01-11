// Simple Golang reference API using Chi router with basic GET and POST endpoints.
//
// Author: SCOTT R. HENZ
// Date: 10/02/2022

package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	// Print message to console
	fmt.Println("Starting Server...")
	// Define router
	router := chi.NewRouter()
	// Define a GET endpoint
	router.Get("/api/getExample", getHandler)
	// Define a POST endpoint
	router.Post("/api/postExample", postHandler)
	// Print message to console
	fmt.Println("Server is listening on port 8088.")
	// Define port and handler
	log.Fatal(http.ListenAndServe(":8088", router))
}

// Function to handle GET endpoint
func getHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Welcome to my GET endpoint.")

	// Service logic here

	// Error handling
	if err != nil {
		return
	}
}

// Function to handle POST endpoint
func postHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode("Welcome to my POST endpoint.")

	// Service logic here

	// Error handling
	if err != nil {
		return
	}
}
