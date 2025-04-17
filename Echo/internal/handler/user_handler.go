package handler

import (
	"User-Management-Go-React/Echo/internal/model"
	"User-Management-Go-React/Echo/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var requestData struct {
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	user := &model.User{
		Username: requestData.Username,
		Email:    requestData.Email,
		Password: requestData.Password,
	}

	err := h.UserService.CreateUser(user, requestData.ConfirmPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, user)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *UserHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	token, err := h.UserService.Login(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (h *UserHandler) GetUserProfile(c echo.Context) error {
	// Get email from JWT claims
	email := c.Get("email").(string)

	// Get user from database
	user, err := h.UserService.GetUserByEmail(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	// Return user profile without sensitive information
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}

func (h *UserHandler) UpdateUserProfile(c echo.Context) error {
	// Get email from JWT claims
	email := c.Get("email").(string)

	// Bind the request data (new profile data)
	var requestData struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	if err := c.Bind(&requestData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// Get the user by email
	user, err := h.UserService.GetUserByEmail(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	// Update user profile
	user.Username = requestData.Username
	user.Email = requestData.Email

	// Save the updated user to the database
	err = h.UserService.UpdateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return the updated user profile
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}

func (h *UserHandler) DeleteUserAccount(c echo.Context) error {
	// Get email from JWT claims
	email := c.Get("email").(string)

	// Delete user from database
	err := h.UserService.DeleteUser(email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Return success message
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Account deleted successfully",
	})
}

func (h *UserHandler) Logout(c echo.Context) error {
	// Since we're using JWT, the actual logout happens on the client side
	// This endpoint is just for consistency and future extensibility
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Logged out successfully",
	})
}
