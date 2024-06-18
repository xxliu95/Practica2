package Generic

import "testing"

func TestAnyOfEmpty(t *testing.T) {
	slice := []int{}
	condition := func(n int) bool { return n > 3 }
	if AnyOf(slice, condition) {
		t.Errorf("Expected false, got true")
	}
}

func TestAnyOfIntegers(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	condition := func(n int) bool { return n > 3 }
	if !AnyOf(slice, condition) {
		t.Errorf("Expected true, got false")
	}
}

func TestAnyOfString(t *testing.T) {
	slice := []string{"Gabriel", "Alberto", "Liu", "JM", "JA", "Gloria"}
	condition := func(s string) bool { return len(s) > 3 }
	if !AnyOf(slice, condition) {
		t.Errorf("Expected true, got false")
	}
}

func TestEqualIntegers(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	slice3 := []int{1, 2, 3, 4}
	slice4 := []int{1, 2, 3, 4, 6}
	if !Equal(slice1, slice2) {
		t.Errorf("Expected true, got false")
	}
	if Equal(slice1, slice3) {
		t.Errorf("Expected false, got true")
	}
	if Equal(slice1, slice4) {
		t.Errorf("Expected false, got true")
	}
}

func TestEqualString(t *testing.T) {
	slice1 := []string{"Gabriel", "Alberto", "Liu", "JM", "JA", "Gloria"}
	slice2 := []string{"Gabriel", "Alberto", "Liu", "JM", "JA", "Gloria"}
	slice3 := []string{"Gabriel", "Alberto", "Liu", "JA", "Gloria"}
	slice4 := []string{"Gabriel", "Albert", "Liu", "JM", "JA", "Gloria"}
	if !Equal(slice1, slice2) {
		t.Errorf("Expected true, got false")
	}
	if Equal(slice1, slice3) {
		t.Errorf("Expected false, got true")
	}
	if Equal(slice1, slice4) {
		t.Errorf("Expected false, got true")
	}
}
