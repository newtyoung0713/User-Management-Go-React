package main

import (
	"context"
	"net/http"
	"time"

	"User-Management-Go-React/Echo/internal/config"
	"User-Management-Go-React/Echo/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize the database (no return value)
	config.InitDB()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusOK, "Hello from Go Echo!")
	})

	// 註冊 POST 路由
	e.POST("/register", func(c echo.Context) error {
		var user model.User

		// 綁定 JSON 資料到 User 結構
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid input",
			})
		}

		// 驗證電子郵件格式
		if !user.ValidateEmail() {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid email format",
			})
		}

		// 密碼加密
		if err := user.HashPassword(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to hash password",
			})
		}

		// 使用資料庫進行插入
		query := `INSERT INTO users (username, email, password, created_at, updated_at) 
		          VALUES ($1, $2, $3, $4, $5)`
		_, err := config.DB.Exec(context.Background(), query, user.Username, user.Email, user.Password, time.Now(), time.Now())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to register user",
			})
		}

		// 返回成功訊息
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Registration successful",
		})
	})

	// 登录 POST 路由
	e.POST("/login", func(c echo.Context) error {
		var req model.User // Get the request from Login page
		// 绑定 JSON 数据到请求结构
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid input",
			})
		}

		// Find user
		var user model.User // Find the user by the value from Login page
		err := config.DB.QueryRow(context.Background(), "SELECT username, email, password FROM users WHERE email=$1", req.Email).Scan(&user.Username, &user.Email, &user.Password)
		if err != nil {
			// 如果用户不存在，返回 401 Unauthorized 错误
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid email or password",
			})
		}

		// 验证密码
		if err := user.CheckPassword(req.Password); err != nil {
			// 如果密码不匹配，返回 401 Unauthorized 错误
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Invalid email or password",
			})
		}

		// 登录成功，可以返回 JWT 或者其他令牌
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Login successful",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
