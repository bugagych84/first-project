package taskService

import (
	"gorm.io/gorm"
)

// Интрефейс репозитория
type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTasks() ([]Task, error)
	GetTaskById(taskId string) (Task, error)
	DeleteTaskById(taskId string) error
	UpdateTask(task Task) error
}

// Структура репозитория
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *taskRepository) GetTaskById(taskId string) (Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", taskId).Error
	return task, err
}

func (r *taskRepository) DeleteTaskById(taskId string) error {
	return r.db.Delete(&Task{}, "id = ?", taskId).Error
}

func (r *taskRepository) UpdateTask(task Task) error {
	return r.db.Save(&task).Error
}
