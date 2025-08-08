package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LucasMartinsVieira/go-todo-api/internal/config"
	"github.com/LucasMartinsVieira/go-todo-api/internal/database"
	db "github.com/LucasMartinsVieira/go-todo-api/internal/database/repository"
	"github.com/LucasMartinsVieira/go-todo-api/internal/todo"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Study Golang", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Work", Completed: false},
}

func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func toggleTodoStatus(c *gin.Context) {
	id := c.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	todo.Completed = !todo.Completed

	c.IndentedJSON(http.StatusOK, todo)
}

func main() {
	cfg := config.LoadConfig()

	pool := database.ConnectDatabase(cfg)
	queries := db.New(pool)

	repo := todo.NewRepository(queries)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	r := gin.Default()
	handler.RegisterRoutes(r)
	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	r.PATCH("/todos/:id", toggleTodoStatus)

	log.Printf("🚀 Server running on %s", addr)
	r.Run(addr)

}
