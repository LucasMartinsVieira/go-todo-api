package todo

import (
	"context"

	db "github.com/LucasMartinsVieira/go-todo-api/internal/database/repository"
)

type Service interface {
	GetTodos(ctx context.Context) ([]db.Todo, error)
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
