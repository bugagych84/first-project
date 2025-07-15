package taskService

import (
	"firstproject/internal/models"
	"fmt"
	"github.com/oapi-codegen/runtime/types"
)

type TaskService interface {
	CreateTask(task models.Task) ([]models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTasksByUserID(userID types.UUID) ([]models.Task, error)
	GetTaskById(taskId string) (models.Task, error)
	UpdateTask(taskId string, newTask models.Task) ([]models.Task, error)
	DeleteTaskById(taskId string) ([]models.Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAllTasks()
}

func (s taskService) GetTasksByUserID(userID types.UUID) ([]models.Task, error) {
	tasks, err := s.repo.GetTasksByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks for user: %w", err)
	}
	return tasks, nil
}

func (s taskService) CreateTask(task models.Task) ([]models.Task, error) {
	if task.UserID == (types.UUID{}) {
		return nil, fmt.Errorf("user ID is required")
	}

	if err := s.repo.CreateTask(task); err != nil {
		return nil, err
	}

	return s.GetAllTasks()
}

func (s taskService) GetTaskById(taskId string) (models.Task, error) {
	task, err := s.repo.GetTaskById(taskId)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s taskService) UpdateTask(taskId string, newTask models.Task) ([]models.Task, error) {
	task, err := s.repo.GetTaskById(taskId)
	if err != nil {
		return nil, err
	}

	if newTask.Name != "" {
		task.Name = newTask.Name
	}
	if newTask.IsDone != nil {
		task.IsDone = newTask.IsDone
	}
	if newTask.UserID != (types.UUID{}) {
		task.UserID = newTask.UserID
	}

	if err := s.repo.UpdateTask(task); err != nil {
		return nil, err
	}

	return s.GetAllTasks()
}

func (s taskService) DeleteTaskById(taskId string) ([]models.Task, error) {
	if err := s.repo.DeleteTaskById(taskId); err != nil {
		return nil, err
	}
	return s.GetAllTasks()
}
