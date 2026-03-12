package usecases_test

import (
	"go.clean/internal/infrastructure"
	"go.clean/internal/usecases"
)

func newUC() *usecases.TodoUseCase {
	return usecases.NewUseCase(infrastructure.NewMemoryRepo())
}
