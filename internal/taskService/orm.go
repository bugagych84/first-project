package taskService

// Task структура сущности задачи
type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"done"`
}
