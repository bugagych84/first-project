package taskService

import openapi_types "github.com/oapi-codegen/runtime/types"

// Task структура сущности задачи
type Task struct {
	ID     openapi_types.UUID `gorm:"primaryKey" json:"id"`
	Name   string             `json:"name"`
	IsDone *bool              `json:"is_done"`
}
