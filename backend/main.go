package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gcox32/query-builder/backend/db"
	"github.com/gcox32/query-builder/backend/routes"
)

func main() {
	// Initialize database
	db.InitDB()

	// Create Fiber app
	app := fiber.New()

	// Register routes
	routes.SetupUserRoutes(app)

	// Start the server
	log.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}