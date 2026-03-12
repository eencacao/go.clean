package adapters

import (
	"encoding/json"
	"net/http"

	"go.clean/internal/usecases"
)

type Handler struct {
	uc *usecases.TodoUseCase
}

type createBody struct {
	Title string `json:"title"`
}

func NewHandler(uc *usecases.TodoUseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) ListAll(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.uc.GetAll())
}

func (h *Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid id")
		return
	}
	todo, ok := h.uc.GetByID(id)
	if !ok {
		writeErr(w, http.StatusNotFound, "not found")
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var b createBody
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid body")
		return
	}
	if b.Title == "" {
		writeErr(w, http.StatusBadRequest, "title required")
		return
	}
	writeJSON(w, http.StatusCreated, h.uc.Create(b.Title))
}
