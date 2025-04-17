package repository

import (
	"User-Management-Go-React/Echo/internal/model"
	"time"

	"github.com/google/uuid"
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
	err := r.DB.Where("deleted_at IS NULL").Find(&users).Error
	return users, err
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	err := r.DB.Model(&model.User{}).
		Where("email = ? AND deleted_at IS NULL", user.Email).
		Updates(user).Error
	return user, err
}

func (r *UserRepository) DeleteUser(email string) error {
	return r.DB.Where("email = ?", email).Delete(&model.User{}).Error
}

// Soft delete
func (r *UserRepository) SoftDeleteUser(userID uuid.UUID) error {
	now := time.Now()
	err := r.DB.Model(&model.User{}).
		Where("id = ? AND deleted_at IS NULL", userID).
		Update("deleted_at", &now).Error
	return err
}

// Restoring a soft deleted user
func (r *UserRepository) RestoreUser(userID uuid.UUID) error {
	return r.DB.Model(&model.User{}).
		Where("id = ? AND deleted_at IS NOT NULL", userID).
		Update("deleted_at", nil).Error
}

// Query undeleted users by ID
func (r *UserRepository) FindByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.DB.Where(&user, "id = ? AND deleted_at IS NULL", id).First(&user).Error
	return &user, err
}

// Query undeleted users by Email
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	return &user, err
}

// Query soft-deleted users
func (r *UserRepository) FindDeletedUser() ([]model.User, error) {
	var users []model.User
	err := r.DB.Where("deleted_at IS NOT NULL").Find(&users).Error
	return users, err
}
