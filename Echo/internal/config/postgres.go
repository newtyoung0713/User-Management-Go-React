package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB global database connection instance
var DB *sql.DB

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

	// Open the database connection
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Test the database connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Unable to Ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
}
