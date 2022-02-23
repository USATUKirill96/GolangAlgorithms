package medium

import (
	"USATUKirill96/leetcode/helpers"
	"fmt"
	"testing"
)

func TestMoveElementToEnd(t *testing.T) {

	tests := []struct {
		arr    []int
		target int
		result []int
	}{
		{
			arr:    []int{2, 1, 2, 2, 2, 3, 4, 2},
			target: 2,
			result: []int{4, 1, 3, 2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		res := MoveElementToEnd(tt.arr, tt.target)
		if !helpers.CompareIntegerLists(res, tt.result) {
			t.Errorf(fmt.Sprint(res))
		}
	}
}

func MoveElementToEnd(arr []int, target int) []int {
	var L, R int
	R = len(arr) - 1

	for L < R {
		if arr[R] == target {
			R -= 1
			continue
		}

		if arr[L] == target {
			arr[L], arr[R] = arr[R], arr[L]
		}

		L += 1
	}

	return arr
}
