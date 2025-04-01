package data

type PriorityLevel string

// Priority levels for todos.
const (
	PriorityLow    PriorityLevel = "low"
	PriorityMedium PriorityLevel = "medium"
	PriorityHigh   PriorityLevel = "high"
)

// Todo represents a task in the todo list.
type Todo struct {
	ID          int64         `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Owner       string        `json:"owner"`
	Priority    PriorityLevel `json:"priority"`
	Completed   bool          `json:"completed"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	DueDate     string        `json:"due_date"`
}

// Todoer is an interface that defines the methods for managing todos.
type Todoer interface {
	Create(todo Todo) (int64, error)
	GetAll() ([]Todo, error)
	GetByID(id int64) (Todo, error)
	Update(todo Todo) error
	Delete(id int64) error
}

// TodoModel is a concrete implementation of the Todoer interface.
// It uses an in-memory slice to store todos.
type TodoModel struct {
	todos []Todo
}

// NewTodoModel creates a new TodoModel instance.
// It initializes the todos slice to an empty slice.
func NewTodoModel() TodoModel {
	return TodoModel{
		todos: []Todo{},
	}
}

// Create adds a new todo to the model.
func (m *TodoModel) Create(todo Todo) (int64, error) {
	// Simulate creating a todo by appending it to the slice
	todo.ID = int64(len(m.todos) + 1)
	m.todos = append(m.todos, todo)
	return todo.ID, nil
}

// GetAll retrieves all todos from the model.
func (m *TodoModel) GetAll() ([]Todo, error) {
	// Simulate getting all todos
	return m.todos, nil
}

// GetByID retrieves a todo by its ID from the model.
func (m *TodoModel) GetByID(id int64) (Todo, error) {
	// Simulate getting a todo by ID
	for _, todo := range m.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, nil
}

// Update modifies an existing todo in the model.
func (m *TodoModel) Update(todo Todo) error {
	// Simulate updating a todo
	for i, t := range m.todos {
		if t.ID == todo.ID {
			m.todos[i] = todo
			return nil
		}
	}
	return nil
}

// Delete removes a todo by its ID from the model.
func (m *TodoModel) Delete(id int64) error {
	// Simulate deleting a todo
	for i, t := range m.todos {
		if t.ID == id {
			m.todos = append(m.todos[:i], m.todos[i+1:]...)
			return nil
		}
	}
	return nil
}
