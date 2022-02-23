package medium

import (
	"fmt"
	"testing"
)

func TestLongestPeak(t *testing.T) {

	tests := []struct {
		input  []int
		result int
	}{
		{
			input:  []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3},
			result: 6,
		},
	}

	for _, tt := range tests {
		res := LongestPeak(tt.input)
		if res != tt.result {
			fmt.Printf("Error in LongestPeak: expected %v, got %v", tt.result, res)
			t.Errorf(fmt.Sprint(res))
		}
	}
}

func LongestPeak(arr []int) int {
	var lp int
	i := 1
	for i < len(arr)-1 {
		var length, right int
		length = 1
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			for j := i - 1; j > 0 && arr[j] < arr[j+1]; j-- {
				length++
			}
			for j := i + 1; j < len(arr) && arr[j] < arr[j-1]; j++ {
				length++
				right = j
			}
			if length > lp {
				lp = length
			}

			i = right
			continue
		} else {
			i++
		}
	}
	return lp
}
