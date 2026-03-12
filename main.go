package main

import (
	"log"
	"net/http"

	"go.clean/internal/adapters"
	"go.clean/internal/infrastructure"
	"go.clean/internal/usecases"
)

func setupRoutes(mux *http.ServeMux, h *adapters.Handler) {
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.ListAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})
	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetOne(w, r)
		case http.MethodPut:
			h.Update(w, r)
		case http.MethodDelete:
			h.Delete(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})
}

func main() {
	repo := infrastructure.NewMemoryRepo()
	uc := usecases.NewUseCase(repo)
	h := adapters.NewHandler(uc)
	mux := http.NewServeMux()
	setupRoutes(mux, h)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
