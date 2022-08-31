package app

import (
	"asrofiw/belajar-golang-restful-api/controller"
	"asrofiw/belajar-golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/v1/categories", categoryController.FindAll)
	router.GET("/v1/categories/:categoryId", categoryController.FindById)
	router.POST("/v1/categories", categoryController.Create)
	router.PUT("/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/v1/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
