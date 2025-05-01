package main

import (
	"go-fiber/src/config"
	"go-fiber/src/database/seeders"
	"go-fiber/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get port from .env (default to 3000 if not set)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // Default port
	}

	if len(os.Args) > 1 && os.Args[1] == "seed" {
		seeders.Run()
		log.Println("Database seeding.")
		return
	} else if len(os.Args) > 1 && os.Args[1] == "migrate" {
		config.ConnectDB()
		return
	}

	// Initialize database and run migrations
	config.ConnectDB()

	// Setup Fiber app
	app := fiber.New()
	routes.SetupRoutes(app)
	routes.SetupAuthRoutes(app) // Add auth routes

	// Start server with dynamic port
	log.Printf("Server is running on port %s", port)
	app.Listen(":" + port)
}
