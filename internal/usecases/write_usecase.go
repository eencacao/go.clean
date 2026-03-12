package usecases

import "go.clean/internal/entities"

func (u *TodoUseCase) Update(
	id int, title string, done bool,
) (*entities.Todo, bool) {
	return u.repo.Update(id, title, done)
}

func (u *TodoUseCase) Delete(id int) bool {
	return u.repo.Delete(id)
}
