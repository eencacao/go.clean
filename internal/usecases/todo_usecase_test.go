package usecases_test

import "testing"

func TestCreate(t *testing.T) {
	uc := newUC()
	todo := uc.Create("buy milk")
	if todo.ID != 1 {
		t.Fatalf("want id=1 got %d", todo.ID)
	}
	if todo.Title != "buy milk" {
		t.Fatalf("want title=%q got %q", "buy milk", todo.Title)
	}
	if todo.Completed {
		t.Fatal("want completed=false")
	}
}

func TestGetAll(t *testing.T) {
	uc := newUC()
	if len(uc.GetAll()) != 0 {
		t.Fatal("want empty list")
	}
	uc.Create("a")
	uc.Create("b")
	if len(uc.GetAll()) != 2 {
		t.Fatalf("want 2 got %d", len(uc.GetAll()))
	}
}

func TestGetByID(t *testing.T) {
	uc := newUC()
	uc.Create("x")
	got, ok := uc.GetByID(1)
	if !ok || got.Title != "x" {
		t.Fatal("expected todo with title x")
	}
	_, ok = uc.GetByID(99)
	if ok {
		t.Fatal("expected not found for id 99")
	}
}

func TestUpdate(t *testing.T) {
	uc := newUC()
	uc.Create("old title")
	got, ok := uc.Update(1, "new title", true)
	if !ok || got.Title != "new title" || !got.Completed {
		t.Fatal("update failed or wrong values")
	}
	_, ok = uc.Update(99, "x", false)
	if ok {
		t.Fatal("expected not found for missing id")
	}
}
