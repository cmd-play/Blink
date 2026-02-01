package main

import (
	"log"

	"blink-server/server"
)

func main() {
	// Load configuration from application.json
	err := server.LoadConfig("application.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Load companies from companies.json
	err = server.LoadCompanies("companies.json")
	if err != nil {
		log.Fatalf("Failed to load companies: %v", err)
	}

	// Start the server
	server.StartServer()
}
