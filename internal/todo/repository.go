package todo

import (
	"context"

	db "github.com/LucasMartinsVieira/go-todo-api/internal/database/repository"
)

type Repository interface {
	ListTodos(ctx context.Context) ([]db.Todo, error)
	CreateTodo(ctx context.Context, arg db.InsertTodoParams) (TodoModel, error)
	GetTodoById(ctx context.Context, id int32) (TodoModel, error)
	ToggleTodoStatusById(ctx context.Context, id int32) (TodoModel, error)
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

func (r *repository) CreateTodo(ctx context.Context, todo db.InsertTodoParams) (TodoModel, error) {
	created, err := r.q.InsertTodo(ctx, db.InsertTodoParams{
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
	})

	if err != nil {
		return TodoModel{}, err
	}

	return mapDBTodo(created), nil
}

func (r *repository) GetTodoById(ctx context.Context, id int32) (TodoModel, error) {
	todo, err := r.q.FindTodoById(ctx, id)

	if err != nil {
		return TodoModel{}, err
	}

	return mapDBTodo(todo), nil

}

func (r *repository) ToggleTodoStatusById(ctx context.Context, id int32) (TodoModel, error) {
	todo, err := r.q.ToggleTodoStatus(ctx, id)

	if err != nil {
		return TodoModel{}, err
	}

	return mapDBTodo(todo), nil
}

func mapDBTodo(dbTodo db.Todo) TodoModel {
	return TodoModel{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: dbTodo.Description.String,
		Status:      dbTodo.Status,
	}
}
