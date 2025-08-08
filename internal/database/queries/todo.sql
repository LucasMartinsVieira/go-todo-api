-- name: FindAllTodos :many
SELECT * from todos;

-- name: InsertTodo :one
INSERT INTO todos (title, description, status) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: FindTodoById :one
SELECT id, title, description, status, created_at, updated_at FROM todos
WHERE id = $1;
