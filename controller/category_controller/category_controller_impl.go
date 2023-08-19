package categorycontroller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	categoryweb "belajar-golang-restful-api/model/web/category_web"
	categoryservice "belajar-golang-restful-api/service/category_service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService categoryservice.CategoryService
}

func NewCategoryController(categoryService categoryservice.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// read and create the request body
	categoryRequest := categoryweb.CategoryRequest{}
	helper.ReadFromRequestBody(r, &categoryRequest)

	//	create response body
	categoryResponse := controller.CategoryService.Create(r.Context(), categoryRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponse,
	}

	//	write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// read and create the request body
	categoryUpdateRequest := categoryweb.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	// find paramater by id
	categoryUpdateRequest.CategoryId = p.ByName("categoryId")

	//	create response body
	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponse,
	}

	//	write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// find paramater by id
	categoryId := p.ByName("categoryId")

	//	create response body
	controller.CategoryService.Delete(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
	}

	//	write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// find paramater by id
	categoryId := p.ByName("categoryId")

	//	create response body
	categoryResponse := controller.CategoryService.FindById(r.Context(), categoryId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponse,
	}

	//	write response
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//	create response body
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponses,
	}

	//	write response
	helper.WriteToResponseBody(w, webResponse)
}
