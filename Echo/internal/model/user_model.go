package model

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username      string     `gorm:"unique;not null;size:50"`
	Email         string     `gorm:"unique;not null;size:100"`
	Password      string     `gorm:"not null;size:255"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime"`
	LastLoginAt   time.Time  `gorm:"default:null"`
	LoginAttempts int        `gorm:"default:0"`
	IsLocked      bool       `gorm:"default:false"`
	AvatarURL     string     `gorm:"size:255;default:null"`
	PreferredMFA  string     `gorm:"size:50;default:null"`
	DeletedAt     *time.Time `gorm:"index" json:"deleted_at"`
}

type PasswordStrengthLevel string

const (
	PasswordWeak   PasswordStrengthLevel = "Weak"
	PasswordMedium PasswordStrengthLevel = "Medium"
	PasswordStrong PasswordStrengthLevel = "Strong"
)

// ValidateEmail validates the email format
func (user *User) ValidateEmail() bool {
	regex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return regex.MatchString(user.Email)
}

// EvaluatePasswordStrength returns the strength level (weak / medium / strong)
func (user *User) EvaluatePasswordStrength() PasswordStrengthLevel {
	// Minimum 8 characters and must contain uppercase letters, lowercase letters, and numbers
	password := user.Password

	if len(password) < 8 {
		return PasswordWeak
	}

	var hasUpper, hasLower, hasNumber, hasSymbol bool

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case strings.ContainsRune("!@#$%^&*()", char):
			hasSymbol = true
		}
	}

	score := 0
	if hasUpper {
		score++
	}
	if hasLower {
		score++
	}
	if hasNumber {
		score++
	}
	if hasSymbol {
		score++
	}

	switch {
	case score >= 3:
		return PasswordStrong
	case score == 2:
		return PasswordMedium
	default:
		return PasswordWeak
	}
}

// Check if the password is a bcrypt hash
func isBcryptHash(s string) bool {
	// bcrypt hashes start with $2a$, $2b$ or $2y$
	return len(s) == 60 && (s[:4] == "$2a$" || s[:4] == "$2b$" || s[:4] == "$2y$")
}

// HashPassword hashes the user's password using bcrypt
func (user *User) HashPassword() error {
	// If the password is already encrypted, do not encrypt it again
	if !isBcryptHash(user.Password) { // bcrypt hash passwords are usually 60 characters long
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil
}

// CheckPassword compares a plain password with the hashed password stored in the database.
func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}
