package infrastructure

import (
	"sync"

	"go.clean/internal/entities"
)

// MemoryRepo stores todos in memory.
type MemoryRepo struct {
	todos   map[int]entities.Todo
	counter int
	mu      sync.RWMutex
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{todos: make(map[int]entities.Todo)}
}

func (r *MemoryRepo) GetAll() []entities.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]entities.Todo, 0, len(r.todos))
	for _, t := range r.todos {
		list = append(list, t)
	}
	return list
}

func (r *MemoryRepo) GetByID(id int) (*entities.Todo, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	t, ok := r.todos[id]
	if !ok {
		return nil, false
	}
	return &t, true
}

func (r *MemoryRepo) Save(todo entities.Todo) entities.Todo {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.counter++
	todo.ID = r.counter
	r.todos[r.counter] = todo
	return todo
}
