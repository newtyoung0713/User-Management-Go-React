package repository

import (
	"User-Management-Go-React/Echo/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := r.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.First(&user, id).Error
	return &user, err
}
