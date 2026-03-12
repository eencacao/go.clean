package adapters_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.clean/internal/adapters"
	"go.clean/internal/infrastructure"
	"go.clean/internal/usecases"
)

func newServer() http.Handler {
	repo := infrastructure.NewMemoryRepo()
	uc := usecases.NewUseCase(repo)
	return adapters.NewServer(uc)
}

func TestAPIGetAll(t *testing.T) {
	srv := newServer()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200 got %d", w.Code)
	}
	var res []map[string]any
	json.NewDecoder(w.Body).Decode(&res)
	if len(res) != 0 {
		t.Fatal("want empty array")
	}
}

func TestAPICreate(t *testing.T) {
	srv := newServer()
	body := bytes.NewBufferString(`{"title":"test todo"}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("want 201 got %d", w.Code)
	}
	var res map[string]any
	json.NewDecoder(w.Body).Decode(&res)
	if res["title"] != "test todo" {
		t.Fatalf("want title=%q got %v", "test todo", res["title"])
	}
}

func TestAPIGetByID(t *testing.T) {
	srv := newServer()
	body := bytes.NewBufferString(`{"title":"find me"}`)
	req := httptest.NewRequest(http.MethodPost, "/todos", body)
	req.Header.Set("Content-Type", "application/json")
	srv.ServeHTTP(httptest.NewRecorder(), req)
	req2 := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
	w2 := httptest.NewRecorder()
	srv.ServeHTTP(w2, req2)
	if w2.Code != http.StatusOK {
		t.Fatalf("want 200 got %d", w2.Code)
	}
}
