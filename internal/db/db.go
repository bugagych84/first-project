package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Указатель на объект подключения к базе данных
var db *gorm.DB

// InitDB функция подключения к базе данных
func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, nil
}
