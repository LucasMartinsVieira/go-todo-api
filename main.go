package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id,omitempty"`
	Item      string `json:"item,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}

var todos = []Todo{
	{ID: "1", Item: "Study Golang", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Work", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)

	router.Run() // listen and serve on 0.0.0.0:8080

}
