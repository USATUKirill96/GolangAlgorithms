package main

import "testing"

func TestRiverSizes(t *testing.T) {
	riverMap := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}
	RiverSizes(riverMap)
}

func RiverSizes(riverMap [][]int) []int {

	var visitedNodes [][]bool
	for range riverMap {
		visitedNodes = append(visitedNodes, make([]bool, len(riverMap[0])))
	}

	var visit func(int, int) int
	visit = func(row, column int) int {
		if visitedNodes[row][column] == true {
			return 0
		}
		if riverMap[row][column] != 1 {
			visitedNodes[row][column] = true
			return 0
		}
		visitedNodes[row][column] = true
		length := 1
		if row > 0 {
			length += visit(row-1, column)
		}
		if column > 0 {
			length += visit(row, column-1)
		}
		if row < len(riverMap)-1 {
			length += visit(row+1, column)
		}
		if column < len(riverMap[0])-1 {
			length += visit(row, column+1)
		}
		return length
	}

	var riverLengths []int
	for row, rowValue := range riverMap {
		for column := range rowValue {
			riverLength := visit(row, column)
			if riverLength != 0 {
				riverLengths = append(riverLengths, riverLength)
			}
		}
	}
	return riverLengths
}
