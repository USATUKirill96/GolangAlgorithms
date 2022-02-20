package main

import "testing"

func TestSameBST(t *testing.T) {
	input1 := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
	input2 := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}
	res := SameBST(input1, input2)
	if !res {
		t.Error("Error in TestSameBST")
	}
}

func SameBST(first, second []int) bool {
	if len(first) == 0 && len(second) == 0 {
		return true
	}
	if len(first) != len(second) {
		return false
	}
	if first[0] != second[0] {
		return false
	}

	root := first[0]
	smallerFirst, greaterFirst := split(first, root)
	smallerSecond, greaterSecond := split(second, root)

	return SameBST(smallerFirst, smallerSecond) == true && SameBST(greaterFirst, greaterSecond)
}

func split(input []int, root int) ([]int, []int) {
	var greater, smaller []int
	for _, v := range input {
		switch {
		case v == root:
			continue
		case v < root:
			smaller = append(smaller, v)
		default:
			greater = append(greater, v)
		}
	}
	return smaller, greater
}
