package config

import (
	"net/http"

	"github.com/rizky-ds/belajar-go-clean-arch/entity"
	"github.com/rizky-ds/belajar-go-clean-arch/helper"
	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func NewHttpRouter() *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = ErrorHandler
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		helper.WriteToResponseBody(writer, entity.HttpResponse{
			Code: http.StatusNotFound,
			Data: "Not found",
		})
	})
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		helper.WriteToResponseBody(writer, entity.HttpResponse{
			Code: http.StatusMethodNotAllowed,
			Data: "Method not allowed",
		})
	})

	return router
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}
	if validationErrors(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		httpResponse := entity.HttpResponse{
			Code: http.StatusNotFound,
			Data: exception.Error,
		}

		helper.WriteToResponseBody(writer, httpResponse)
		return true
	} else {
		return false
	}
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		httpResponse := entity.HttpResponse{
			Code: http.StatusBadRequest,
			Data: exception.Error(),
		}

		helper.WriteToResponseBody(writer, httpResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	httpResponse := entity.HttpResponse{
		Code: http.StatusInternalServerError,
		Data: err,
	}

	helper.WriteToResponseBody(writer, httpResponse)
}
