package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestLevenshteinDistance(t *testing.T) {
	str1 := "abc"
	str2 := "yabd"
	res := LevenshteinDistance(str1, str2)
	if res != 2 {
		t.Errorf("Expected: %v, got: %v", 2, res)
	}
}

func LevenshteinDistance(str1, str2 string) int {

	// We form a table of minimum actions to transform origin string to target
	//    | "" | y | a | b | d |
	// "" |  0 | 1 | 2 | 3 | 4 |
	// a  |  1 | 1 | 1 | 2 | 3 |
	// b  | 2  | 2 | 2 | 1 | 2 |
	// c  | 3  | 3 | 3 | 2 | 2 |
	//--------------------------

	// transform inputs into list of rows, columns headers
	origin := strings.Split(str1, "")
	origin = append([]string{" "}, origin...)

	target := strings.Split(str2, "")
	target = append([]string{" "}, target...)

	// Init the "ways" table with zeros
	var ways [][]int
	for i := 0; i <= len(str1); i++ {
		ways = append(ways, make([]int, len(str2)+1))
	}

	// generic function to iterate through rows or columns of the "ways" table. It takes initial values
	// where iterate from and a function increasing the row or column value
	iterate := func(row, column int, inc func(int, int) (int, int)) {
		for ; row < len(origin) && column < len(target); column, row = inc(column, row) {
			if origin[row] == target[column] { // If the letters are equal
				if row == 0 && column == 0 { // edge case of [0; 0] coordinate
					continue
				}
				ways[row][column] = ways[row-1][column-1] // replace with diagonal value
			} else {
				ways[row][column] = getMinNeighbor(column, row, ways) + 1 // replace with minimal neighbor + 1
			}
		}
	}

	for row, column := 0, 0; row < len(origin) && column < len(target); column, row = column+1, row+1 {
		iterate(row, column, func(a, b int) (int, int) { return a, b + 1 }) // iterate through columns
		iterate(row, column, func(a, b int) (int, int) { return a + 1, b }) // iterate through rows
	}

	fmt.Println(ways)

	return ways[len(ways)-1][len(ways[0])-1]
}

func getMinNeighbor(column, row int, table [][]int) int {
	min := math.MaxInt
	getMin := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	if column > 0 {
		min = getMin(min, table[row][column-1])
	}
	if row > 0 {
		min = getMin(min, table[row-1][column])
	}
	if row > 0 && column > 0 {
		min = getMin(min, table[row-1][column-1])
	}
	if min == math.MaxInt {
		min = 0
	}
	return min
}
