package handlers

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"

	"firstproject/internal/models"
	"firstproject/internal/taskService"
	"firstproject/internal/web/tasks"
	"github.com/labstack/echo/v4"
	types "github.com/oapi-codegen/runtime/types"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

// GetTasks returns all tasks
func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := make(tasks.GetTasks200JSONResponse, 0, len(allTasks))
	for _, tsk := range allTasks {
		id := tsk.ID
		userId := tsk.UserID
		response = append(response, tasks.Task{
			Id:     &id,
			UserId: userId,
			Name:   tsk.Name,
			IsDone: tsk.IsDone,
		})
	}

	return response, nil
}

// PostTasks creates a new task
func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body is required")
	}

	taskToCreate := models.Task{
		ID:     types.UUID(uuid.New()),
		UserID: request.Body.UserId,
		Name:   request.Body.Name,
		IsDone: request.Body.IsDone,
	}

	createdTasks, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := make(tasks.PostTasks201JSONResponse, 0, len(createdTasks))
	for _, t := range createdTasks {
		id := t.ID
		userId := t.UserID
		response = append(response, tasks.Task{
			Id:     &id,
			UserId: userId,
			Name:   t.Name,
			IsDone: t.IsDone,
		})
	}

	return response, nil
}

// DeleteTasksId deletes a task by ID
func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id.String()

	remainingTasks, err := h.service.DeleteTaskById(taskID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to delete task: %w", err).Error())
	}

	response := make(tasks.DeleteTasksId200JSONResponse, 0, len(remainingTasks))
	for _, t := range remainingTasks {
		id := t.ID
		userId := t.UserID
		response = append(response, tasks.Task{
			Id:     &id,
			UserId: userId,
			Name:   t.Name,
			IsDone: t.IsDone,
		})
	}

	return response, nil
}

// PatchTasksId updates a task by ID
func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body required")
	}

	updateData := models.Task{
		ID:     request.Id,
		UserID: request.Body.UserId,
		Name:   request.Body.Name,
		IsDone: request.Body.IsDone,
	}

	updatedTasks, err := h.service.UpdateTask(request.Id.String(), updateData)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := make(tasks.PatchTasksId200JSONResponse, 0, len(updatedTasks))
	for _, t := range updatedTasks {
		id := t.ID
		userId := t.UserID
		response = append(response, tasks.Task{
			Id:     &id,
			UserId: userId,
			Name:   t.Name,
			IsDone: t.IsDone,
		})
	}

	return response, nil
}
