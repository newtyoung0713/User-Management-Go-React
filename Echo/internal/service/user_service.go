package service

import (
	"User-Management-Go-React/Echo/internal/model"
	"User-Management-Go-React/Echo/internal/repository"
	"errors"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.UserRepo.GetAllUsers()
}

func (s *UserService) CreateUser(user *model.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required fields")
	}

	// user.Password = hashPassword(user.Password)
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	var user model.User

	// Assume that a database operation is called here, such as GORM's First method
	return &user, nil
}
