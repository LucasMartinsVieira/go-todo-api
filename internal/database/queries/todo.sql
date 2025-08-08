-- name: FindAllTodos :many
SELECT * from todos;

-- name: InsertTodo :one
INSERT INTO todos (title, description, status) 
VALUES ($1, $2, $3)
RETURNING *;
