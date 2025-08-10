package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/LucasMartinsVieira/go-todo-api/docs"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.Title = "Gin Todo App"
	docs.SwaggerInfo.Description = "A simple gin todo application"
	docs.SwaggerInfo.Version = "1.0"

	r.GET("/todos", h.getTodos)
	r.POST("/todos", h.createTodo)
	r.GET("/todo/:id", h.getTodo)
	r.PATCH("/todo/:id", h.toggleTodoStatus)
}

// getTodos godoc
// @Summary      List all todos
// @Description  Retrieve all todos from the database
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Success      200  {array}   TodoModel
// @Failure      404  {object}  map[string]string
// @Router       /todos [get]
func (h *Handler) getTodos(c *gin.Context) {
	todos, err := h.service.GetTodos(c.Request.Context())

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error:": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todos)
}

// createTodo godoc
// @Summary 			Create a new todo
// @Description 	Create a new todo in the database
// @Tags 					Todos
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} CreateTodoSchema
// @Failure       404  {object}  map[string]string
// @Router       /todos [post]
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

// getTodo godoc
// @Summary      Get a todo by ID
// @Description  Retrieve a specific todo from the database by its ID
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  TodoModel
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [get]
func (h *Handler) getTodo(c *gin.Context) {
	id64, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	id := int32(id64)

	todo, err := h.service.GetTodo(c.Request.Context(), id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

// toggleTodoStatus godoc
// @Summary      Toggle todo status
// @Description  Toggle the completion status of a todo by its ID
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  TodoModel
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /todo/{id} [patch]
func (h *Handler) toggleTodoStatus(c *gin.Context) {
	var input ToggleTodoStatusSchema

	if err := c.ShouldBindUri(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.service.ToggleTodoStatus(c.Request.Context(), input)

	fmt.Println("TODO", &todo)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}
