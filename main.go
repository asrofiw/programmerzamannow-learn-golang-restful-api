package main

import (
	"asrofiw/belajar-golang-restful-api/app"
	"asrofiw/belajar-golang-restful-api/controller"
	"asrofiw/belajar-golang-restful-api/helper"
	"asrofiw/belajar-golang-restful-api/middleware"
	"asrofiw/belajar-golang-restful-api/repository"
	"asrofiw/belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
