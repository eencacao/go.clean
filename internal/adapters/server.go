package adapters

import (
	"net/http"

	"go.clean/internal/usecases"
)

// NewServer wires up all routes and returns the mux.
func NewServer(uc *usecases.TodoUseCase) *http.ServeMux {
	h := NewHandler(uc)
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(
		w http.ResponseWriter, r *http.Request,
	) {
		switch r.Method {
		case http.MethodGet:
			h.ListAll(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			http.Error(w, "method not allowed", 405)
		}
	})
	mux.HandleFunc("/todos/", func(
		w http.ResponseWriter, r *http.Request,
	) {
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
	return mux
}
