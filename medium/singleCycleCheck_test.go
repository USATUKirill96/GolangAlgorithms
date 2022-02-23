package medium

import (
	"testing"
)

func TestSingleCycleCheck(t *testing.T) {
	input := []int{2, 3, 1, -4, -4, 2}
	res := singleCycleCheck(input)
	if res != true {
		t.Errorf("Expected: %v, got: %v", true, res)
	}
}

func singleCycleCheck(input []int) bool {
	var newPosition func(int) int
	newPosition = func(shift int) int {
		if shift > len(input) {
			shift -= len(input)
			return newPosition(shift)
		}
		if shift < 0 {
			shift = len(input) + shift
			return newPosition(shift)
		}
		return shift
	}
	visited := make([]int, len(input))
	var position int
	for range input {
		visited[position] += 1
		position = newPosition(position + input[position])
	}
	for _, v := range visited {
		if v != 1 {
			return false
		}
	}
	return true
}
