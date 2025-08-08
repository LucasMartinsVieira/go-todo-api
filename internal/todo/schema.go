package todo

type CreateTodoSchema struct {
	Title       string `json:"title" binding: "required"`
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status" binding: "required"`
}
