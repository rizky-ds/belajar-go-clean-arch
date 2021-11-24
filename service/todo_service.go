package service

import (
	"context"

	"github.com/rizky-ds/belajar-go-clean-arch/entity"
)

type TodoService interface {
	Create(ctx context.Context, request entity.TodoCreateRequest) entity.TodoResponse
	List(ctx context.Context) []entity.TodoResponse
	Get(ctx context.Context, id uint64) entity.TodoResponse
	Update(ctx context.Context, request entity.TodoUpdateRequest) entity.TodoResponse
	Delete(ctx context.Context, id uint64)
}
