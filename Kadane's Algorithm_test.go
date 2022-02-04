package main

import (
	"testing"
)

func TestKadaneAlg(t *testing.T) {
	input := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	res := KadaneAlg(input)
	if res != 19 {
		t.Errorf("Expected: %v, got: %v", 19, res)
	}
}

func KadaneAlg(input []int) int {
	var currentMax, globalMax int
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for _, v := range input {
		currentMax = max(currentMax+v, v)
		globalMax = max(globalMax, currentMax)
	}

	return globalMax
}
