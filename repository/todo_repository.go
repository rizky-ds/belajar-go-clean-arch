package repository

import (
	"context"

	"github.com/rizky-ds/belajar-go-clean-arch/entity"
)

type TodoRepository interface {
	Save(ctx context.Context, todo *entity.Todo) error
	FindAll(ctx context.Context) ([]entity.Todo, error)
	FindById(ctx context.Context, id uint64) (entity.Todo, error)
	Update(ctx context.Context, todo *entity.Todo) error
	Delete(ctx context.Context, id uint64) error
}
