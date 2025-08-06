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
}

func (h *Handler) getTodos(c *gin.Context) {
	todos, err := h.service.GetTodos(c.Request.Context())

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todos)
}
