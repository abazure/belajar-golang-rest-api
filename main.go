package main

import (
	"belajar-golang-rest-api/app"
	"belajar-golang-rest-api/controller"
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/repository"
	"belajar-golang-rest-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
