package interfaces

import "go.clean/internal/entities"

// TodoRepository is the contract for todo storage.
type TodoRepository interface {
	GetAll() []entities.Todo
	GetByID(id int) (*entities.Todo, bool)
	Save(todo entities.Todo) entities.Todo
	Update(id int, title string, done bool) (*entities.Todo, bool)
	Delete(id int) bool
}
