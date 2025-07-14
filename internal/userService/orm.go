package userService

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Task структура сущности задачи
type User struct {
	ID       openapi_types.UUID `gorm:"primaryKey" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
