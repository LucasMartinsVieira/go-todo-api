package todo

import (
	"context"

	db "github.com/LucasMartinsVieira/go-todo-api/internal/database/repository"
)

type Repository interface {
	ListTodos(ctx context.Context) ([]db.Todo, error)
	// GetTodo(ctx context.Context, id int32) (db.Todo, error)
	// CreateTodo(ctx context.Context, arg d)
}

type repository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) Repository {
	return &repository{q: q}
}

func (r *repository) ListTodos(ctx context.Context) ([]db.Todo, error) {
	return r.q.FindAllTodos(ctx)
}
