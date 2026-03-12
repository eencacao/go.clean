package adapters

import (
	"encoding/json"
	"net/http"
)

type updateBody struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid id")
		return
	}
	var b updateBody
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid body")
		return
	}
	todo, ok := h.uc.Update(id, b.Title, b.Completed)
	if !ok {
		writeErr(w, http.StatusNotFound, "not found")
		return
	}
	writeJSON(w, http.StatusOK, todo)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		writeErr(w, http.StatusBadRequest, "invalid id")
		return
	}
	if !h.uc.Delete(id) {
		writeErr(w, http.StatusNotFound, "not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
