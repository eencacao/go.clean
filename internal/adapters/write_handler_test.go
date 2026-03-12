package adapters_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIUpdate(t *testing.T) {
	srv := newServer()
	post := bytes.NewBufferString(`{"title":"original"}`)
	r := httptest.NewRequest(http.MethodPost, "/todos", post)
	r.Header.Set("Content-Type", "application/json")
	srv.ServeHTTP(httptest.NewRecorder(), r)
	upd := bytes.NewBufferString(`{"title":"updated","completed":true}`)
	r2 := httptest.NewRequest(http.MethodPut, "/todos/1", upd)
	r2.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r2)
	if w.Code != http.StatusOK {
		t.Fatalf("want 200 got %d", w.Code)
	}
}

func TestAPIDelete(t *testing.T) {
	srv := newServer()
	post := bytes.NewBufferString(`{"title":"bye"}`)
	r := httptest.NewRequest(http.MethodPost, "/todos", post)
	r.Header.Set("Content-Type", "application/json")
	srv.ServeHTTP(httptest.NewRecorder(), r)
	r2 := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r2)
	if w.Code != http.StatusNoContent {
		t.Fatalf("want 204 got %d", w.Code)
	}
}

func TestAPINotFound(t *testing.T) {
	srv := newServer()
	r := httptest.NewRequest(http.MethodGet, "/todos/99", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	if w.Code != http.StatusNotFound {
		t.Fatalf("want 404 got %d", w.Code)
	}
}
