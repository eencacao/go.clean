package infrastructure

import "go.clean/internal/entities"

func (r *MemoryRepo) Update(
	id int, title string, done bool,
) (*entities.Todo, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.todos[id]
	if !ok {
		return nil, false
	}
	t.Title = title
	t.Completed = done
	r.todos[id] = t
	return &t, true
}

func (r *MemoryRepo) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.todos[id]; !ok {
		return false
	}
	delete(r.todos, id)
	return true
}
