package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gcox32/query-builder/backend/db"
	"github.com/gcox32/query-builder/backend/routes"
)

func main() {
	// Initialize database
	db.InitDB()

	// Create Fiber app
	app := fiber.New()

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173,http://127.0.0.1:5173",  // SvelteKit dev server
		AllowHeaders: "Origin, Content-Type, Accept, X-Sandbox-Mode",
	}))

	// Register routes
	routes.SetupUserRoutes(app)
	routes.SetupConnectionRoutes(app)

	// Start the server
	log.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}