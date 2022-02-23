package medium

import (
	"sort"
	"testing"
)

func TestMinNumberOfCoins(t *testing.T) {
	val := 6
	coins := []int{1, 2, 4}
	expected := 2

	res := MinNumberOfCoins(val, coins)
	if res != expected {
		t.Errorf("Error in MinNumberOfCoins expected %v, got %v", expected, res)
	}
}

func MinNumberOfCoins(val int, coins []int) int {
	sort.Ints(coins)
	var ways []int
	for i := 0; i <= val; i++ {
		ways = append(ways, 0)
	}

	for _, c := range coins {
		for wi := range ways {
			if c <= wi {
				ways[wi] = ways[wi-c] + 1
			}
		}
	}
	return ways[len(ways)-1]
}
