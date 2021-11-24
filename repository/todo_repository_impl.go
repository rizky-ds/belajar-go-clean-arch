package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/rizky-ds/belajar-go-clean-arch/entity"
)

type TodoRepositoryImpl struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &TodoRepositoryImpl{
		db: db,
	}
}

func (repo *TodoRepositoryImpl) Save(ctx context.Context, todo *entity.Todo) error {
	sql := "insert into todo(title, description) values(?, ?)"
	result, err := repo.db.ExecContext(ctx, sql, todo.Title, todo.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	todo.Id = uint64(id)

	return nil
}

func (repo *TodoRepositoryImpl) FindAll(ctx context.Context) ([]entity.Todo, error) {
	sql := "select id, title, description from todo"
	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []entity.Todo{}
	for rows.Next() {
		todo := entity.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *TodoRepositoryImpl) FindById(ctx context.Context, id uint64) (entity.Todo, error) {
	sql := "select id, title, description from todo where id=?"
	rows, err := repo.db.QueryContext(ctx, sql)
	if err != nil {
		return entity.Todo{}, err
	}
	defer rows.Close()

	todo := entity.Todo{}
	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description)
		if err != nil {
			return entity.Todo{}, err
		} else {
			return todo, nil
		}
	}

	msg := fmt.Sprintf("Todo with Id %d not found", id)
	return todo, errors.New(msg)
}

func (repo *TodoRepositoryImpl) Update(ctx context.Context, todo *entity.Todo) error {
	sql := "update todo set title=?, description=? where id=?"
	_, err := repo.db.ExecContext(ctx, sql, todo.Title, todo.Description, todo.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *TodoRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	sql := "delete from todo where id=?"
	_, err := repo.db.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}
