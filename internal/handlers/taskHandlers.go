package handlers

import (
	"firstproject/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, tasks)
}

// Основные методы работы с задачами
func (h *TaskHandler) PostTask(c echo.Context) error {
	var req taskService.Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	tasks, err := h.service.CreateTask(req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, tasks)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	tasks, err := h.service.DeleteTaskById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) PatchTask(c echo.Context) error {
	id := c.Param("id")
	var req taskService.Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	tasks, err := h.service.UpdateTask(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tasks)
}
