package todo

import (
	"net/http"
	"strconv"

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
	r.GET("/todo/:id", h.getTodo)
}

func (h *Handler) getTodos(c *gin.Context) {
	todos, err := h.service.GetTodos(c.Request.Context())

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error:": err.Error()})
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

func (h *Handler) getTodo(c *gin.Context) {
	id64, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	id := int32(id64)

	todo, err := h.service.GetTodoById(c.Request.Context(), id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}
