package model

import (
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null;size:50"`
	Email     string    `gorm:"unique;not null;size:100"`
	Password  string    `gorm:"not null;size:255"`
	IsDeleted bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// ValidateEmail validates the email format
func (u *User) ValidateEmail() bool {
	regex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return regex.MatchString(u.Email)
}

// HashPassword hashes the user's password using bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares a plain password with the hashed password stored in the database.
func (u *User) CheckPassword(password string) error {
	// bcrypt.CompareHashAndPassword compares the given password with the hashed password
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
