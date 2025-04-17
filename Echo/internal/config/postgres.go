package config

import (
	"User-Management-Go-React/Echo/internal/model"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	// Configure GORM logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore record not found error
			Colorful:                  true,        // Enable color
		},
	)

	// Configure GORM
	config := &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true, // Cache prepared statements
	}

	// Establish a connection
	var errGorm error
	DB, errGorm = gorm.Open(postgres.Open(dsn), config)
	if errGorm != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)           // Maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)          // Maximum number of open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum amount of time a connection may be reused
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)

	// Auto Migrate the schema
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Log the success message
	log.Println("Successfully connected to the database and migrated the schema")
}
