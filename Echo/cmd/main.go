package main

import (
	"log"
	"net/http"

	"User-Management-Go-React/Echo/internal/config"
	"User-Management-Go-React/Echo/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database (no return value)
	config.InitDB()

	// Use the DB instance that has been initialized in the config package
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: config.DB, // Use the *sql.DB from config
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Only run AutoMigrate if table doesn't exist
	var tableExists bool
	db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users')").Scan(&tableExists)

	if !tableExists {
		err = db.AutoMigrate(&model.User{})
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/users", func(c echo.Context) error {
		var users []model.User
		result := db.Limit(5).Find(&users)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		}
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		user := new(model.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		result := db.Create(user)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		}
		return c.JSON(http.StatusCreated, user)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
