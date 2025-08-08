package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/todos", h.getTodos)
	r.POST("/todos", h.createTodo)
}

func (h *Handler) getTodos(c *gin.Context) {
	todos, err := h.service.GetTodos(c.Request.Context())

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func (h *Handler) createTodo(c *gin.Context) {
	var input CreateTodoSchema

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.CreateTodo(c.Request.Context(), input)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, todo)
}
