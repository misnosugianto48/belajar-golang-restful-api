package routes

import (
	categorycontroller "belajar-golang-restful-api/controller/category_controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController categorycontroller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/todos/", categoryController.FindAll)
	router.GET("/api/todos/:todoId", categoryController.FindById)
	router.POST("/api/todos", categoryController.Create)
	router.PUT("/api/todos/:todoId", categoryController.Update)
	router.DELETE("/api/todos/:todoId", categoryController.Delete)

	return router
}
