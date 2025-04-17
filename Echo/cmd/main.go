package main

import (
	"User-Management-Go-React/Echo/internal/config"
	"User-Management-Go-React/Echo/internal/handler"
	"User-Management-Go-React/Echo/internal/middleware"
	"User-Management-Go-React/Echo/internal/repository"
	"User-Management-Go-React/Echo/internal/service"
	"log"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize the database (no return value)
	config.InitDB()

	// Echo instance
	e := echo.New()

	// Add middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())

	// Initialize dependencies
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Public routes
	e.POST("/api/users", userHandler.CreateUser)
	e.POST("/api/users/login", userHandler.Login)

	// Protected routes
	api := e.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/users", userHandler.GetUsers)
		api.GET("/users/profile", userHandler.GetUserProfile)
	}

	// Start the server
	log.Fatal(e.Start(":1323"))
}
