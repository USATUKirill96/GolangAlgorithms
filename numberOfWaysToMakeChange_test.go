package main

import (
	"sort"
	"testing"
)

func TestNumberOfWaysToMakeChange(t *testing.T) {
	val := 10
	coins := []int{1, 5, 10, 25}
	expected := 4

	res := NumberOfWaysToMakeChange(val, coins)
	if res != expected {
		t.Errorf("Error in numberOfWaysToMakeChange: expected %v, got %v", expected, res)
	}
}

func NumberOfWaysToMakeChange(val int, coins []int) int {
	sort.Ints(coins)
	var ways []int
	for i := 0; i <= val; i++ {
		ways = append(ways, 0)
	}
	ways[0] = 1

	for _, v := range coins {
		for i := range ways {
			if v <= i {
				ways[i] += ways[i-v]
			}
		}
	}
	return ways[len(ways)-1]
}
