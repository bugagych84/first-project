package handlers

import (
	"context"
	"fmt"
	"net/http"

	"firstproject/internal/taskService"
	"firstproject/internal/web/tasks"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Name:   tsk.Name,
			IsDone: tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body is required")
	}

	newUUID := openapi_types.UUID(uuid.New())

	taskToCreate := taskService.Task{
		ID:     newUUID,
		Name:   request.Body.Name,
		IsDone: request.Body.IsDone,
	}

	createdTasks, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := tasks.PostTasks201JSONResponse{}
	for _, t := range createdTasks {
		task := tasks.Task{
			Id:     &t.ID,
			Name:   t.Name,
			IsDone: t.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id.String()

	remainingTasks, err := h.service.DeleteTaskById(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete task: %w", err)
	}

	response := make(tasks.DeleteTasksId200JSONResponse, 0, len(remainingTasks))
	for _, t := range remainingTasks {
		uuidVal := openapi_types.UUID(t.ID)
		task := tasks.Task{
			Id:     &uuidVal,
			Name:   t.Name,
			IsDone: t.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Простая валидация
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body required")
	}

	// Обновляем только переданные поля
	updateData := taskService.Task{
		ID: request.Id,
	}

	if request.Body.Name != "" {
		updateData.Name = request.Body.Name
	}

	if request.Body.IsDone != nil {
		updateData.IsDone = request.Body.IsDone
	}

	updatedTasks, err := h.service.UpdateTask(request.Id.String(), updateData)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response tasks.PatchTasksId200JSONResponse
	for _, t := range updatedTasks {
		idCopy := t.ID
		response = append(response, tasks.Task{
			Id:     &idCopy,
			Name:   t.Name,
			IsDone: t.IsDone,
		})
	}

	return response, nil
}
