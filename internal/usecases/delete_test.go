package usecases_test

import "testing"

func TestDelete(t *testing.T) {
	uc := newUC()
	uc.Create("to delete")
	if !uc.Delete(1) {
		t.Fatal("expected delete ok")
	}
	if uc.Delete(1) {
		t.Fatal("expected delete fail for missing")
	}
}
