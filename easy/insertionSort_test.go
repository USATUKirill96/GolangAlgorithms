package easy

import (
	"USATUKirill96/leetcode/helpers"
	"testing"
)

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

func TestInsertionSort(t *testing.T) {

	tests := []struct {
		input  []int
		result []int
	}{
		{
			input:  []int{2, 7, 11, 15},
			result: []int{2, 7, 11, 15},
		},
		{
			input:  []int{3, 2, 4},
			result: []int{2, 3, 4},
		},
		{
			input:  []int{10, 0, 4, 2, 1, 3, 3},
			result: []int{0, 1, 2, 3, 3, 4, 10},
		},
	}

	for _, tt := range tests {
		res := insertionSort(tt.input)
		if !helpers.CompareIntegerLists(res, tt.result) {
			t.Errorf("Error in insertionSort function. Expected %v, got %v", tt.result, res)
		}
	}
}

func insertionSort(arr []int) []int {
	for i := range arr {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j -= 1
		}
	}
	return arr
}
