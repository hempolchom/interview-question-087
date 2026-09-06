package main

import (
	"fmt"
	"log"

	"example.com/interview-question-087/go-backend/config"
	"example.com/interview-question-087/go-backend/database"
	"example.com/interview-question-087/go-backend/routes"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDB()

	r := routes.SetupRouter()

	port := cfg.Port
	if port == "" {
		port = "9087"
	}

	log.Printf("Server is starting on port %s...", port)
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}