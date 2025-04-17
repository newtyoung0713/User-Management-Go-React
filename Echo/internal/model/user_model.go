package model

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username      string    `gorm:"unique;not null;size:50"`
	Email         string    `gorm:"unique;not null;size:100"`
	Password      string    `gorm:"not null;size:255"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	LastLoginAt   time.Time `gorm:"default:null"`
	LoginAttempts int       `gorm:"default:0"`
	IsLocked      bool      `gorm:"default:false"`
	AvatarURL     string    `gorm:"size:255;default:null"`
	PreferredMFA  string    `gorm:"size:50;default:null"`
	IsDeleted     bool      `gorm:"default:false"`
}

// ValidateEmail validates the email format
func (user *User) ValidateEmail() bool {
	regex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return regex.MatchString(user.Email)
}

// HashPassword hashes the user's password using bcrypt
func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares a plain password with the hashed password stored in the database.
func (user *User) CheckPassword(password string) error {
	// bcrypt.CompareHashAndPassword compares the given password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
