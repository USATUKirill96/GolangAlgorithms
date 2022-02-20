package main

import "testing"

func TestMaxSumIncreasingSubsequences(t *testing.T) {
	input := []int{8, 12, 2, 3, 15, 5, 7}
	expected := 35

	result := MaxSumIncreasingSubsequence(input)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func MaxSumIncreasingSubsequence(input []int) int {
	sums := make([]int, len(input))
	var result int

	for i, v := range input {
		var maxSum int
		for j := 0; j < i; j++ {
			if input[j] < v && sums[j] > maxSum {
				maxSum = sums[j]
			}
		}
		sums[i] = maxSum + v
		if sums[i] > result {
			result = sums[i]
		}
	}
	return result
}
