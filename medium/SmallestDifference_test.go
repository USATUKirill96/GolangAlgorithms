package medium

import (
	"USATUKirill96/leetcode/helpers"
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestSmallestDifference(t *testing.T) {

	tests := []struct {
		first  []int
		second []int
		result []int
	}{
		{
			first:  []int{-1, 5, 10, 20, 28, 3},
			second: []int{26, 134, 135, 15, 17},
			result: []int{28, 26},
		},
	}

	for _, tt := range tests {
		res := SmallestDifference(tt.first, tt.second)
		if !helpers.CompareIntegerLists(res, tt.result) {
			t.Errorf(fmt.Sprint(res))
		}
	}
}

func SmallestDifference(first, second []int) []int {
	sort.Ints(first[:])
	sort.Ints(second[:])

	var result struct {
		pair []int
		diff int
	}
	result.diff = math.MaxInt

	abs := func(n int) int {
		if n >= 0 {
			return n
		} else {
			return -n
		}
	}

	var L, R int

	for L < len(first) && R < len(second) {
		if first[L] == second[R] {
			return []int{first[L], second[R]}
		}

		diff := abs(first[L] - second[R])
		if result.diff > diff {
			result.diff = diff
			result.pair = []int{first[L], second[R]}
		}
		if first[L] < second[R] {
			L += 1
		} else {
			R += 1
		}
	}
	return result.pair
}
