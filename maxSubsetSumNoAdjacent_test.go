package main

import (
	"math"
	"testing"
)

func TestMaxSubset(t *testing.T) {
	arr := []int{7, 10, 12, 7, 9, 14}
	res := MaxSubset(arr)
	if res != 33 {
		t.Errorf("Expected: %v, got: %v", 33, res)
	}

}

func MaxSubset(arr []int) int {
	max := func(a, b int) int {
		return int(math.Max(
			float64(a),
			float64(b),
		))
	}
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) < 3 {
		return max(arr[0], arr[1])
	}
	result := []int{arr[0], arr[1]}
	for i := 2; i < len(arr); i++ {
		result = append(
			result,
			max(result[i-1], result[i-2]+arr[i]),
		)
	}
	return result[len(result)-1]
}
