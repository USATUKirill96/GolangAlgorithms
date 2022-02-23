package medium

import (
	"fmt"
	"testing"
)

func TestSearchInSortedMatrix(t *testing.T) {
	input := [][]int{
		{1, 4, 7, 12, 15, 1000},
		{2, 5, 19, 31, 32, 1001},
		{3, 8, 24, 33, 35, 1002},
		{40, 41, 42, 44, 45, 1003},
		{99, 100, 13, 106, 128, 1004},
	}

	target := 44

	res := SearchInSortedMatrix(input, target)
	fmt.Println(res)
}

func SearchInSortedMatrix(input [][]int, target int) []int {
	// Top right corner coordinates
	row := 0
	column := len(input[0]) - 1

	for {
		if input[row][column] == target {
			return []int{row, column}
		}
		if input[row][column] > target {
			column -= 1
		} else {
			row += 1
		}
	}

}
