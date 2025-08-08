package todo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/LucasMartinsVieira/go-todo-api/internal/database/repository"
)

type Service interface {
	GetTodos(ctx context.Context) ([]db.Todo, error)
	CreateTodo(ctx context.Context, input CreateTodoSchema) (TodoModel, error)
	GetTodo(ctx context.Context, id int32) (TodoModel, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetTodos(ctx context.Context) ([]db.Todo, error) {
	return s.repo.ListTodos(ctx)
}

func (s *service) CreateTodo(ctx context.Context, input CreateTodoSchema) (TodoModel, error) {
	// TODO: Change this pgtype.Text to another thing
	todo := db.InsertTodoParams{
		Title: input.Title,
		Description: pgtype.Text{
			String: input.Description,
			Valid:  input.Description != "",
		},
		Status: input.Status,
	}

	return s.repo.CreateTodo(ctx, todo)
}

func (s *service) GetTodo(ctx context.Context, id int32) (TodoModel, error) {
	return s.repo.GetTodo(ctx, id)
}
