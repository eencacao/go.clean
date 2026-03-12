package usecases

import (
	"go.clean/internal/entities"
	"go.clean/internal/interfaces"
)

// TodoUseCase holds all business logic.
type TodoUseCase struct {
	repo interfaces.TodoRepository
}

func NewUseCase(r interfaces.TodoRepository) *TodoUseCase {
	return &TodoUseCase{repo: r}
}

func (u *TodoUseCase) GetAll() []entities.Todo {
	return u.repo.GetAll()
}

func (u *TodoUseCase) GetByID(id int) (*entities.Todo, bool) {
	return u.repo.GetByID(id)
}

func (u *TodoUseCase) Create(title string) entities.Todo {
	return u.repo.Save(entities.NewTodo(0, title))
}
