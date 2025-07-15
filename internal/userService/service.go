package userService

import (
	"firstproject/internal/models"
	"firstproject/internal/taskService"
	"fmt"
	types "github.com/oapi-codegen/runtime/types"
)

type UserService interface {
	CreateUser(user models.User) ([]models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(userId types.UUID) (models.User, error)
	UpdateUser(userId types.UUID, newUser models.User) ([]models.User, error)
	DeleteUserById(userId types.UUID) ([]models.User, error)
	GetTasksForUser(userID types.UUID) ([]models.Task, error)
}

type userService struct {
	repo        UserRepository
	taskService taskService.TaskService
}

func NewUserService(r UserRepository, ts taskService.TaskService) UserService {
	return &userService{
		repo:        r,
		taskService: ts,
	}
}

func (s userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s userService) CreateUser(user models.User) ([]models.User, error) {
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}
	return s.GetAllUsers()
}

func (s userService) GetUserById(userId types.UUID) (models.User, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s userService) UpdateUser(userId types.UUID, newUser models.User) ([]models.User, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	// Update only allowed fields
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Password != "" {
		user.Password = newUser.Password
	}

	if err := s.repo.UpdateUser(user); err != nil {
		return nil, err
	}

	return s.GetAllUsers()
}

func (s userService) DeleteUserById(userId types.UUID) ([]models.User, error) {
	// First check if user has tasks
	tasks, err := s.taskService.GetTasksByUserID(userId)
	if err == nil && len(tasks) > 0 {
		return nil, fmt.Errorf("cannot delete user with existing tasks")
	}

	if err := s.repo.DeleteUserById(userId); err != nil {
		return nil, err
	}

	return s.GetAllUsers()
}

func (s userService) GetTasksForUser(userID types.UUID) ([]models.Task, error) {
	tasks, err := s.taskService.GetTasksByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks for user: %w", err)
	}
	return tasks, nil
}
