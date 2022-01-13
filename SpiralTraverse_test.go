package main

import (
	"USATUKirill96/leetcode/helpers"
	"fmt"
	"testing"
)

func TestSpiralTraverse(t *testing.T) {

	tests := []struct {
		input  [][]int
		result []int
	}{
		{
			input: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
			result: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		},
	}

	for _, tt := range tests {
		res := SpiralTraverse(tt.input)
		if !helpers.CompareIntegerLists(res, tt.result) {
			fmt.Printf("Error in SpiralTraverse: expected %v, got %v", tt.result, res)
			t.Errorf(fmt.Sprint(res))
		}
	}
}

type coords struct {
	SC int
	EC int
	SR int
	ER int
}

func SpiralTraverse(spiral [][]int) []int {
	c := coords{SC: 0, EC: len(spiral[0]) - 1, SR: 0, ER: len(spiral) - 1}
	var res []int
	for c.EC > c.SC {
		for i := c.SC; i <= c.EC; i++ {
			res = append(res, spiral[c.SR][i])
		}
		for i := c.SR + 1; i <= c.ER; i++ {
			res = append(res, spiral[i][c.EC])
		}
		for i := c.EC - 1; i >= c.SC; i-- {
			res = append(res, spiral[c.ER][i])
		}
		for i := c.ER - 1; i > c.SR; i-- {
			res = append(res, spiral[i][c.SC])
		}
		c.SC++
		c.EC--
		c.SR++
		c.ER--
	}
	return res
}
