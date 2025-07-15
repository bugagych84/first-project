package userService

import (
	"firstproject/internal/models"
	types "github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
)

// Интрефейс репозитория
type UserRepository interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(userId types.UUID) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUserById(userId types.UUID) error
}

// Структура репозитория
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) GetUserById(userId types.UUID) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", userId).Error
	return user, err
}

func (r *userRepository) DeleteUserById(userId types.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", userId).Error
}

func (r *userRepository) UpdateUser(user models.User) error {
	return r.db.Save(&user).Error
}
