package main

import (
	"USATUKirill96/leetcode/helpers"
	"testing"
)

//Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//You can return the answer in any order.

func TestTwoSum(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			result: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			result: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			result: []int{0, 1},
		},
	}

	for _, tt := range tests {
		res := twoSum(tt.nums, tt.target)
		if !helpers.CompareIntegerLists(res, tt.result) {
			t.Errorf("Error in 1.TwoSum function. Expected %v, got %v", tt.result, res)
		}
	}
}

func twoSum(nums []int, target int) []int {
	// we set the map capacity to decrease the memory usage
	m := make(map[int]int, len(nums))

	for i, vi := range nums {
		val := target - vi
		res, found := m[val]
		if found {
			return []int{res, i}
		}
		m[vi] = i
	}

	return []int{}
}
