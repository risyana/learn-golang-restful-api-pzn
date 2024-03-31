package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	newCategoryRepository := repository.NewCategoryRepository()
	newCategoryService := service.NewCategoryService(newCategoryRepository, db, validate)
	newCategoryController := controller.NewCategoryController((newCategoryService))

	router := httprouter.New()

	router.GET("/api/categories", newCategoryController.FindAll)
	router.POST("/api/categories", newCategoryController.Create)
	router.PUT("/api/categories/:categoryId", newCategoryController.Update)
	router.GET("/api/categories/:categoryId", newCategoryController.FindById)
	router.DELETE("/api/categories/:categoryId", newCategoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	println("server started")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
