package userService

import (
	"gorm.io/gorm"
)

// Интрефейс репозитория
type UserRepository interface {
	CreateUser(user User) error
	GetAllUsers() ([]User, error)
	GetUserById(userId string) (User, error)
	DeleteUserById(userId string) error
	UpdateUser(user User) error
}

// Структура репозитория
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) GetUserById(userId string) (User, error) {
	var user User
	err := r.db.First(&user, "id = ?", userId).Error
	return user, err
}

func (r *userRepository) DeleteUserById(userId string) error {
	return r.db.Delete(&User{}, "id = ?", userId).Error
}

func (r *userRepository) UpdateUser(user User) error {
	return r.db.Save(&user).Error
}
