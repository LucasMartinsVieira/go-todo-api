package todo

type TodoModel struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status"`
}
