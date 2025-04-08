package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB global database connection instance
var DB *sql.DB

// Environment variable keys
const (
	EnvHost     = "PG_SQL_HOST"
	EnvPort     = "PG_SQL_PORT"
	EnvUser     = "PG_SQL_USER"
	EnvPassword = "PG_SQL_PASSWORD"
	EnvDBName   = "PG_SQL_NAME"
)

// InitDB initializes the database connection
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read the PostgreSQL configuration from environment variables
	host := os.Getenv(EnvHost)
	port := os.Getenv(EnvPort)
	user := os.Getenv(EnvUser)
	password := os.Getenv(EnvPassword)
	dbname := os.Getenv(EnvDBName)

	// Check if any required environment variable is missing
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatalf("One or more required environment variables are missing: %s, %s, %s, %s, %s",
			EnvHost, EnvPort, EnvUser, EnvPassword, EnvDBName)
	}

	// Build the PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		// Read the database configuration from the environment variable
		host, port, user, password, dbname)

	// Open the database connection
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Test database connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Unable to Ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
}
