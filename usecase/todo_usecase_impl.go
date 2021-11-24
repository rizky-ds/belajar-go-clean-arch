package usecase

import (
	"net/http"
	"strconv"

	"github.com/rizky-ds/belajar-go-clean-arch/entity"
	"github.com/rizky-ds/belajar-go-clean-arch/helper"
	"github.com/rizky-ds/belajar-go-clean-arch/service"
	"github.com/julienschmidt/httprouter"
)

type TodoUsecaseImpl struct {
	todoService service.TodoService
}

func NewTodoUsecase(todoService service.TodoService) TodoUsecase {
	return &TodoUsecaseImpl{
		todoService: todoService,
	}
}

func (usecase *TodoUsecaseImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoCreateRequest := entity.TodoCreateRequest{}
	helper.ReadFromRequestBody(request, &todoCreateRequest)

	todo := usecase.todoService.Create(request.Context(), todoCreateRequest)
	httpResponse := entity.HttpResponse{
		Code: 201,
		Data: todo,
	}

	helper.WriteToResponseBody(writer, httpResponse)
}

func (usecase *TodoUsecaseImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoResponse := usecase.todoService.List(request.Context())
	httpResponse := entity.HttpResponse{
		Code: 200,
		Data: todoResponse,
	}

	helper.WriteToResponseBody(writer, httpResponse)
}

func (usecase *TodoUsecaseImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	todo := usecase.todoService.Get(request.Context(), uint64(id))
	httpResponse := entity.HttpResponse{
		Code: 200,
		Data: todo,
	}

	helper.WriteToResponseBody(writer, httpResponse)
}

func (usecase *TodoUsecaseImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoUpdateRequest := entity.TodoUpdateRequest{}
	helper.ReadFromRequestBody(request, &todoUpdateRequest)

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	todoUpdateRequest.Id = uint64(id)

	todo := usecase.todoService.Update(request.Context(), todoUpdateRequest)
	httpResponse := entity.HttpResponse{
		Code: 200,
		Data: todo,
	}

	helper.WriteToResponseBody(writer, httpResponse)
}

func (usecase *TodoUsecaseImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	usecase.todoService.Delete(request.Context(), uint64(id))
	httpResponse := entity.HttpResponse{
		Code: 200,
	}

	helper.WriteToResponseBody(writer, httpResponse)
}
