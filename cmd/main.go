package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LucasMartinsVieira/go-todo-api/internal/config"
	"github.com/LucasMartinsVieira/go-todo-api/internal/database"
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

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func getTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}

	c.IndentedJSON(http.StatusOK, todo)
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

	fmt.Println(pool)

	router := gin.Default()
	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)

	log.Printf("🚀 Server running on %s", addr)
	router.Run(addr)
}
