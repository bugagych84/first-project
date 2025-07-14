package handlers

import (
	"context"
	"fmt"
	"net/http"

	"firstproject/internal/userService"
	"firstproject/internal/web/users"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type UserHandler struct {
	service userService.UserService
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		user := users.User{
			Id:       &user.ID,
			Email:    user.Email,
			Password: user.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func NewUserHandler(s userService.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body is required")
	}

	newUUID := openapi_types.UUID(uuid.New())

	userToCreate := userService.User{
		ID:       newUUID,
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	createdUsers, err := h.service.CreateUser(userToCreate)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := users.PostUsers201JSONResponse{}
	for _, t := range createdUsers {
		user := users.User{
			Id:       &t.ID,
			Email:    t.Email,
			Password: t.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID := request.Id.String()

	remainingUsers, err := h.service.DeleteUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %w", err)
	}

	response := make(users.DeleteUsersId200JSONResponse, 0, len(remainingUsers))
	for _, t := range remainingUsers {
		uuidVal := openapi_types.UUID(t.ID)
		user := users.User{
			Id:       &uuidVal,
			Email:    t.Email,
			Password: t.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	// Простая валидация
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body required")
	}

	// Обновляем только переданные поля
	updateData := userService.User{
		ID: request.Id,
	}

	if request.Body.Email != "" {
		updateData.Email = request.Body.Email
	}

	updatedUsers, err := h.service.UpdateUser(request.Id.String(), updateData)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response users.PatchUsersId200JSONResponse
	for _, t := range updatedUsers {
		idCopy := t.ID
		response = append(response, users.User{
			Id:       &idCopy,
			Email:    t.Email,
			Password: t.Password,
		})
	}

	return response, nil
}
