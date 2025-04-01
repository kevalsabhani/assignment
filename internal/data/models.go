package data

// Models holds the application models.
type Models struct {
	Todos TodoModel
}

// NewModels initializes a new Models struct with a TodoModel instance.
func NewModels() Models {
	return Models{
		Todos: NewTodoModel(),
	}
}
