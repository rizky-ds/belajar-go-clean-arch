package config

import (
	"database/sql"

	"github.com/rizky-ds/belajar-go-clean-arch/repository"
	"github.com/rizky-ds/belajar-go-clean-arch/service"
	"github.com/rizky-ds/belajar-go-clean-arch/usecase"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func ApplyTodoRoutes(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	todoRepository := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepository, validate)
	todoUsecase := usecase.NewTodoUsecase(todoService)

	router.POST("/api/todos", todoUsecase.Create)
	router.GET("/api/todos", todoUsecase.FindAll)
	router.GET("/api/todos/:id", todoUsecase.FindById)
	router.PUT("/api/todos/:id", todoUsecase.Update)
	router.DELETE("/api/todos/:id", todoUsecase.Delete)
}
