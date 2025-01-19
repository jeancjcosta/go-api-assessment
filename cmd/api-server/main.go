package main

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-api-assessment/configs"
	_ "go-api-assessment/docs"
	"go-api-assessment/handlers"
	"go-api-assessment/utils"
)

func main() {
	// Load the configuration settings
	cfg := configs.LoadConfig()

	// Initialize the logger with the specified log level
	utils.InitLogger(cfg.LogLevel)

	// Load the input file
	utils.LoadFile("input.txt")

	// Create a new router
	router := mux.NewRouter()

	// Register the SearchHandler for the endpoint with a value parameter
	router.HandleFunc("/endpoint/{value}", handlers.SearchHandler).Methods("GET")

	// Serve the Swagger documentation at the /docs/ path
	router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// Log the server start message with the configured port
	log.Printf("Starting server on port %s", cfg.Port)

	// Start the HTTP server on the specified port and log any fatal errors
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
