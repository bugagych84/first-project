package models

import "github.com/oapi-codegen/runtime/types"

type Task struct {
	ID     types.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name   string     `json:"name"`
	IsDone *bool      `json:"is_done"`
	UserID types.UUID `gorm:"type:uuid"`
}
