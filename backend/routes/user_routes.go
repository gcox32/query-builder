package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gcox32/query-builder/backend/db"
	"github.com/gcox32/query-builder/backend/models"
)

// SetupUserRoutes initializes user-related routes
func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	userGroup.Get("/", GetUsers)
	userGroup.Post("/", CreateUser)
}

// Get all users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(users)
}

// Create a user
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	db.DB.Create(&user)
	return c.JSON(user)
}