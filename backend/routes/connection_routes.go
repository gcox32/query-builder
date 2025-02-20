package routes

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"           // PostgreSQL driver
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionTimeout = 5 * time.Second
)

var mockConnections = []Connection{
	{
		ID:       "mock-postgres",
		Name:     "Mock PostgreSQL",
		Type:     "postgresql",
		Host:     "localhost",
		Port:     "5432",
		Database: "mock_db",
		Username: "mock_user",
	},
	{
		ID:       "mock-mysql",
		Name:     "Mock MySQL",
		Type:     "mysql",
		Host:     "localhost",
		Port:     "3306",
		Database: "mock_db",
		Username: "mock_user",
	},
	{
		ID:       "mock-mongodb",
		Name:     "Mock MongoDB",
		Type:     "mongodb",
		Host:     "localhost",
		Port:     "27017",
		Database: "mock_db",
		Username: "mock_user",
	},
}

type Connection struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type TestConnectionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

// SetupConnectionRoutes initializes connection-related routes
func SetupConnectionRoutes(app *fiber.App) {
	connectionGroup := app.Group("/api/connections")
	connectionGroup.Get("/", GetConnections)
	connectionGroup.Post("/test", TestConnection)
}

// GetConnections returns all available connections
func GetConnections(c *fiber.Ctx) error {
	if c.Get("X-Sandbox-Mode") == "true" {
		return c.JSON(mockConnections)
	}
	// TODO: Implement actual connection storage/retrieval
	return c.JSON([]Connection{})
}

// TestConnection attempts to establish a database connection
func TestConnection(c *fiber.Ctx) error {
	conn := new(Connection)
	if err := c.BodyParser(conn); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	start := time.Now()
	
	var success bool
	var message string
	
	if c.Get("X-Sandbox-Mode") == "true" {
		success = true
		message = fmt.Sprintf("Mock connection successful to %s database", conn.Type)
	} else {
		success, message = tryConnection(conn)
	}
	
	elapsed := time.Since(start)

	return c.JSON(TestConnectionResponse{
		Success: success,
		Message: message,
		Time:    fmt.Sprintf("%.2fs", elapsed.Seconds()),
	})
}

func tryConnection(conn *Connection) (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch conn.Type {
	case "postgresql":
		return tryPostgreSQL(ctx, conn)
	case "mysql":
		return tryMySQL(ctx, conn)
	case "mongodb":
		return tryMongoDB(ctx, conn)
	default:
		return false, "Unsupported database type"
	}
}

func tryPostgreSQL(ctx context.Context, conn *Connection) (bool, string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		conn.Host, conn.Port, conn.Username, conn.Password, conn.Database)
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return false, fmt.Sprintf("Failed to create connection: %v", err)
	}
	defer db.Close()

	// Set connection timeout
	db.SetConnMaxLifetime(connectionTimeout)

	err = db.PingContext(ctx)
	if err != nil {
		return false, fmt.Sprintf("Failed to connect: %v", err)
	}

	return true, "Successfully connected to PostgreSQL database"
}

func tryMySQL(ctx context.Context, conn *Connection) (bool, string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=5s",
		conn.Username, conn.Password, conn.Host, conn.Port, conn.Database)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return false, fmt.Sprintf("Failed to create connection: %v", err)
	}
	defer db.Close()

	// Set connection timeout
	db.SetConnMaxLifetime(connectionTimeout)

	err = db.PingContext(ctx)
	if err != nil {
		return false, fmt.Sprintf("Failed to connect: %v", err)
	}

	return true, "Successfully connected to MySQL database"
}

func tryMongoDB(ctx context.Context, conn *Connection) (bool, string) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		conn.Username, conn.Password, conn.Host, conn.Port, conn.Database)
	
	clientOptions := options.Client().
		ApplyURI(uri).
		SetConnectTimeout(connectionTimeout).
		SetServerSelectionTimeout(connectionTimeout)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return false, fmt.Sprintf("Failed to create connection: %v", err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		return false, fmt.Sprintf("Failed to connect: %v", err)
	}

	return true, "Successfully connected to MongoDB database"
} 