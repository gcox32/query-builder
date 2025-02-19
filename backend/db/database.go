package db

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/gcox32/query-builder/backend/models"
)

// DB global variable
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected!")

	// Auto-migrate models
	DB.AutoMigrate(&models.User{})
	log.Println("Database migrated!")
}
