package service

import (
	"context"

	"github.com/rizky-ds/belajar-go-clean-arch/entity"
	"github.com/rizky-ds/belajar-go-clean-arch/repository"
	"github.com/go-playground/validator"
)

type TodoServiceImpl struct {
	todoRepo repository.TodoRepository
	validate *validator.Validate
}

func NewTodoService(todoRepo repository.TodoRepository, validate *validator.Validate) TodoService {
	return &TodoServiceImpl{
		todoRepo: todoRepo,
		validate: validate,
	}
}

func (service *TodoServiceImpl) Create(ctx context.Context, request entity.TodoCreateRequest) entity.TodoResponse {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	todo := entity.Todo{
		Title:       request.Title,
		Description: request.Description,
	}

	err = service.todoRepo.Save(ctx, &todo)
	if err != nil {
		panic(err)
	}

	return entity.ToTodoResponse(&todo)
}

func (service *TodoServiceImpl) List(ctx context.Context) []entity.TodoResponse {
	todos, err := service.todoRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	result := []entity.TodoResponse{}
	for _, todo := range todos {
		result = append(result, entity.ToTodoResponse(&todo))
	}

	return result
}

func (service *TodoServiceImpl) Get(ctx context.Context, id uint64) entity.TodoResponse {
	todo, err := service.todoRepo.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	return entity.ToTodoResponse(&todo)
}

func (service *TodoServiceImpl) Update(ctx context.Context, request entity.TodoUpdateRequest) entity.TodoResponse {
	err := service.validate.Struct(request)
	if err != nil {
		panic(err)
	}

	todo := entity.Todo{
		Id:          request.Id,
		Title:       request.Title,
		Description: request.Description,
	}

	err = service.todoRepo.Update(ctx, &todo)
	if err != nil {
		panic(err)
	}

	return entity.ToTodoResponse(&todo)
}

func (service *TodoServiceImpl) Delete(ctx context.Context, id uint64) {
	err := service.todoRepo.Delete(ctx, id)
	if err != nil {
		panic(err)
	}
}
