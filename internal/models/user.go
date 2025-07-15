package models

import "github.com/oapi-codegen/runtime/types"

type User struct {
	ID       types.UUID `gorm:"primaryKey" json:"id"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Tasks    []Task
}
