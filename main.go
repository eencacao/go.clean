package main

import (
	"log"
	"net/http"

	"go.clean/internal/adapters"
	"go.clean/internal/infrastructure"
	"go.clean/internal/usecases"
)

func main() {
	repo := infrastructure.NewMemoryRepo()
	uc := usecases.NewUseCase(repo)
	mux := adapters.NewServer(uc)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
