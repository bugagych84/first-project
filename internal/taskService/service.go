package taskService

import "github.com/google/uuid"

// Интерфейс сервиса задач
type TaskService interface {
	CreateTask(name string) ([]Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(taskId string) (Task, error)
	UpdateTask(taskId string, newTask Task) ([]Task, error)
	DeleteTaskById(taskId string) ([]Task, error)
}

// Структура сервиса задач
type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s taskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s taskService) CreateTask(name string) ([]Task, error) {
	task := Task{
		ID:     uuid.NewString(),
		Name:   name,
		IsDone: false,
	}

	if err := s.repo.CreateTask(task); err != nil {
		return []Task{}, err
	}

	tasks, err := s.GetAllTasks()
	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}

func (s taskService) GetTaskById(taskId string) (Task, error) {
	task, err := s.repo.GetTaskById(taskId)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (s taskService) UpdateTask(taskId string, newTask Task) ([]Task, error) {
	task, err := s.repo.GetTaskById(taskId)

	if err != nil {
		return []Task{}, err
	}

	task.Name = newTask.Name
	task.IsDone = newTask.IsDone

	if err := s.repo.UpdateTask(task); err != nil {
		return []Task{}, err
	}

	tasks, err := s.GetAllTasks()
	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func (s taskService) DeleteTaskById(taskId string) ([]Task, error) {
	_, err := s.repo.GetTaskById(taskId)

	if err != nil {
		return []Task{}, err
	}

	err = s.repo.DeleteTaskById(taskId)
	if err != nil {
		return []Task{}, err
	}

	tasks, err := s.GetAllTasks()
	if err != nil {
		return []Task{}, err
	}

	return tasks, nil
}
