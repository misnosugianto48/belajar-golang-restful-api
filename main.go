package main

import (
	"belajar-golang-restful-api/app"
	categorycontroller "belajar-golang-restful-api/controller/category_controller"
	"belajar-golang-restful-api/helper"
	categoryrepository "belajar-golang-restful-api/repository/category_repo"
	categoryservice "belajar-golang-restful-api/service/category_service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	db := app.NewDatabase()
	validate := validator.New()
	categoryRepository := categoryrepository.NewCategoryRepository()
	categoryService := categoryservice.NewCategoryService(categoryRepository, db, validate)
	categoryController := categorycontroller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories/", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	log.Printf("Server running on %v", server.Addr)

	err := server.ListenAndServe()
	helper.PanicIfError("", err)
}
