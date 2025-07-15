package main

import (
	"firstproject/internal/userService"
	"firstproject/internal/web/users"
	"log"

	"firstproject/internal/db"
	"firstproject/internal/handlers"
	"firstproject/internal/taskService"
	"firstproject/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dataBase, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	e := echo.New()

	taskRepo := taskService.NewTaskRepository(dataBase)
	serviceTask := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(serviceTask)

	userRepo := userService.NewUserRepository(dataBase)
	serviceUser := userService.NewUserService(userRepo, serviceTask)
	userHandler := handlers.NewUserHandler(serviceUser)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	strictTaskHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, strictTaskHandler)

	strictUserHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	err = e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
}
