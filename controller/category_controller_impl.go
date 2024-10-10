package controller

import (
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/service"
	"belajar-golang-rest-api/web"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImp struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImp{CategoryService: categoryService}
}

func (controller *CategoryControllerImp) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helper.WriteResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImp) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImp) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImp) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImp) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponses,
	}
	helper.WriteResponseBody(writer, webResponse)
}
