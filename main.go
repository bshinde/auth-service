package main

import (
	"log"
	"net/http"

	"auth-service/config"
	"auth-service/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Load configurations
	config.LoadConfig()

	// Create a new router
	router := mux.NewRouter()

	// Register routes
	controllers.RegisterRoutes(router)

	// Start the server
	serverAddress := ":8080"
	log.Printf("Starting server on %s\n", serverAddress)
	err := http.ListenAndServe(serverAddress, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
