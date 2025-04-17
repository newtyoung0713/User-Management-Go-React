package config

import (
	"User-Management-Go-React/Echo/internal/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global database connection instance
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read the PostgreSQL DSN from environment variable
	dsn := os.Getenv("PG_SQL_DSN")
	// Check if the DSN is empty
	if dsn == "" {
		log.Fatalf("PG_SQL_DSN environment variable is not set")
	}

	// Establish a connection
	var errGorm error
	DB, errGorm = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errGorm != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Auto Migrate the schema
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Log the success message
	log.Println("Successfully connected to the database and migrated the schema")
}
