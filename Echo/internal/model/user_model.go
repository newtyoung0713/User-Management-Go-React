package model

import (
	"regexp"
	"time"
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
