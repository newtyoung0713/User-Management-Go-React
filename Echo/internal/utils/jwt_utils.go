package utils

import (
	"User-Management-Go-React/Echo/internal/model"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func getExpirationTime() time.Duration {
	expStr := os.Getenv("JWT_EXPIRATION_TIME")
	if expStr == "" {
		return 12 * time.Hour // Default 12 hours
	}

	// Parse time string (e.g. "12h", "30m", "24h")
	value := expStr[:len(expStr)-1]
	unit := expStr[len(expStr)-1:]

	val, err := strconv.Atoi(value)
	if err != nil {
		return 12 * time.Hour // If parsing fails, return the default value
	}

	switch strings.ToLower(unit) {
	case "h":
		return time.Duration(val) * time.Hour
	case "m":
		return time.Duration(val) * time.Minute
	default:
		return 12 * time.Hour
	}
}

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(getExpirationTime())
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
