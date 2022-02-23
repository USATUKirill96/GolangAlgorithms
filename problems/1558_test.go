package problems

import (
	"testing"
)

func Test1558(t *testing.T) {

	tests := []struct {
		nums []int
		res  int
	}{
		{
			nums: []int{1, 5},
			res:  5,
		},
		{
			nums: []int{2, 2},
			res:  3,
		},
		{
			nums: []int{4, 2, 5},
			res:  6,
		},
		{
			nums: []int{3, 2, 2, 4},
			res:  7,
		},
	}

	for _, tt := range tests {
		res := minOperations(tt.nums)
		if res != tt.res {
			t.Errorf("Error in 1558. Expected %v, got %v", tt.res, res)
		}
	}
}

func minOperations(nums []int) int {
	var counter int

	isNull := func() bool {
		for _, v := range nums {
			if v != 0 {
				return false
			}
		}
		return true
	}

	for {

		for i, v := range nums {
			if v%2 != 0 {
				nums[i] -= 1
				counter += 1
			}
		}

		if isNull() {
			return counter
		}

		for i := 0; i < len(nums); i++ {
			nums[i] = nums[i] / 2
		}

		counter += 1
	}
}
