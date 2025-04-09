package router

import (
	"User-Management-Go-React/Echo/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.GET("/users", userHandler.GetUsers)
	e.POST("/users", userHandler.CreateUser)
}
