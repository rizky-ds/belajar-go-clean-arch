package main

import (
	"net/http"

	"github.com/rizky-ds/belajar-go-clean-arch/config"
	"github.com/rizky-ds/belajar-go-clean-arch/middleware"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.NewDB()
	validate := validator.New()

	router := config.NewHttpRouter()
	config.ApplyTodoRoutes(router, db, validate)
	//config.ApplyArticleRoutes(router, db, validate)
	//config.ApplyXXXRoutes(router, db, validate)

	logMiddleware := middleware.LogMiddleware{HttpHandler: router}

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: &logMiddleware,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
