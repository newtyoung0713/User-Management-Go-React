package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB global database connection instance
var DB *pgx.Conn

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

	// Establish a connection using pgx
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Store the connection globally
	DB = conn
	log.Println("Successfully connected to the database")
}
