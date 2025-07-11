package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var tasks = []Task{}

func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	task := Task{
		ID:   uuid.NewString(),
		Name: req.Name,
		Done: false,
	}
	tasks = append(tasks, task)
	return c.JSON(http.StatusCreated, tasks)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.JSON(http.StatusOK, tasks)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = Task{
				ID:   id,
				Name: req.Name,
				Done: req.Done,
			}

			return c.JSON(http.StatusOK, tasks)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
}

func main() {
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
