package helpers

import (
	"testing"
)

func TestCompareIntegerLists(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := []int{1, 2, 3}
	d := []int{}

	if res := CompareIntegerLists(a, b); res != false {
		t.Errorf("Failed check function. Expected false, got %v ", res)
	}
	if res := CompareIntegerLists(a, c); res != true {
		t.Errorf("Failed check function. Expected true, got %v", res)
	}
	if res := CompareIntegerLists(a, a); res != true {
		t.Errorf("Failed check function. Expected true, got %v", res)
	}
	if res := CompareIntegerLists(a, d); res != false {
		t.Errorf("Failed check function. Expected true, got %v", res)
	}
}
