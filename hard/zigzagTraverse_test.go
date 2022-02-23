package hard

import (
	"fmt"
	"testing"
)

func TestZigZagTraverse(t *testing.T) {
	input := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16},
	}

	result := ZigZagTraverse(input)
	fmt.Println(result)
}

func ZigZagTraverse(input [][]int) []int {
	var row, column int
	direction := "down"
	var result []int
	for row < len(input) && column < len(input[0]) {
		result = append(result, input[row][column])
		if direction == "down" {
			if column == 0 {
				row += 1
				direction = "up"
			} else if row == len(input)-1 {
				column += 1
				direction = "up"
			} else {
				row += 1
				column -= 1
			}
			continue
		}
		if direction == "up" {
			if column == len(input[0])-1 {
				row += 1
				direction = "down"
			} else if row == 0 {
				column += 1
				direction = "down"
			} else {
				row -= 1
				column += 1
			}
			continue
		}
	}
	return result
}
