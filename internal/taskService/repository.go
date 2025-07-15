package taskService

import (
	"firstproject/internal/models"
	"github.com/oapi-codegen/runtime/types"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task models.Task) error
	GetAllTasks() ([]models.Task, error)
	GetTasksByUserID(userID types.UUID) ([]models.Task, error)
	GetTaskById(taskId string) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTaskById(taskId string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task models.Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTasksByUserID(userID types.UUID) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskById(taskId string) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ?", taskId).Error
	return task, err
}

func (r *taskRepository) DeleteTaskById(taskId string) error {
	return r.db.Delete(&models.Task{}, "id = ?", taskId).Error
}

func (r *taskRepository) UpdateTask(task models.Task) error {
	return r.db.Save(&task).Error
}
