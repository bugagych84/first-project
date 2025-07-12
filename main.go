package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Указатель на объект подключения к базе данных
var db *gorm.DB

// Функция подключения к базе данных
func initDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatalf("Failed to migrate tasks: %v", err)
	}
}

// Структура сущности задачи
type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	IsDone bool   `json:"done"`
}

// Хелпер для получения всех тасок
func getAllTasks() ([]Task, error) {
	var tasks []Task
	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

/* Основные методы работы с задачами */
func postTask(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	task := Task{
		ID:     uuid.NewString(),
		Name:   req.Name,
		IsDone: false,
	}

	if err := db.Create(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create task"})
	}

	tasks, err := getAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusCreated, tasks)
}

func getTasks(c echo.Context) error {
	tasks, err := getAllTasks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	if err := db.Delete(&Task{}, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	tasks, err := getAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}
	task = Task{
		ID:     id,
		Name:   req.Name,
		IsDone: req.IsDone,
	}

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update task"})
	}

	tasks, err := getAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

func main() {
	initDB()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/tasks", getTasks)
	e.POST("/tasks", postTask)
	e.DELETE("/tasks/:id", deleteTask)
	e.PATCH("/tasks/:id", patchTask)

	err := e.Start(":8080")

	if err != nil {
		e.Logger.Fatal(err)
	}
}
