package service

import (
	"User-Management-Go-React/Echo/internal/model"
	"User-Management-Go-React/Echo/internal/repository"
	"User-Management-Go-React/Echo/internal/utils"
	"errors"

	"github.com/google/uuid"
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

func (s *UserService) CreateUser(user *model.User, confirmPassword string) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("missing required fields")
	}

	// Validate email format
	if !user.ValidateEmail() {
		return errors.New("invalid email format")
	}

	// Check password strength level
	strength := user.EvaluatePasswordStrength()
	if strength == model.PasswordWeak {
		return errors.New("password is too weak")
	}

	// Check if password and confirmPassword match
	if user.Password != confirmPassword {
		return errors.New("password and confirm password do not match")
	}

	// Hash the password
	if err := user.HashPassword(); err != nil {
		return err
	}

	// Save user in the database
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetUserByID(id uuid.UUID) (*model.User, error) {
	return s.UserRepo.FindByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.UserRepo.FindByEmail(email)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) UpdateUser(user *model.User) error {
	// Necessary logical processing is performed here to update the user information in the database
	_, err := s.UserRepo.UpdateUser(user)
	return err
}

func (s *UserService) DeleteUser(email string) error {
	// First check if user exists
	_, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return err
	}

	// Delete the user
	return s.UserRepo.DeleteUser(email)
}
