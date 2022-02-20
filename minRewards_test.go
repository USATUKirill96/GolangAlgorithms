package main

import (
	"fmt"
	"math"
	"testing"
)

func TestMinRewards(t *testing.T) {
	input := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	res := MinRewardsPeaks(input)
	fmt.Println(res)
}

func MinRewards(input []int) []int {
	rewards := make([]int, len(input))
	for i := range rewards {
		rewards[i] = 1
	}

	for i := range input {
		if i == 0 {
			continue
		}
		if input[i] > input[i-1] {
			rewards[i] = int(math.Max(float64(rewards[i-1]+1), float64(rewards[i])))
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if i == len(input)-1 {
			continue
		}
		if input[i] > input[i+1] {
			rewards[i] = int(math.Max(float64(rewards[i+1]+1), float64(rewards[i])))
		}
	}

	return rewards

}

func MinRewardsPeaks(input []int) []int {
	var pits []int

	rewards := make([]int, len(input))
	for i := range rewards {
		rewards[i] = 1
	}

	get := func(i int) int {
		if i < 0 || i > len(input)-1 {
			return math.MaxInt
		}
		return input[i]
	}

	for i := range input {
		if get(i) < get(i-1) && get(i) < get(i+1) {
			pits = append(pits, i)
		}
	}

	increase := func(point int, prev func(int) int) {
		if point <= 0 || point >= len(input)-1 {
			return
		}
		for i := point; get(i) > get(prev(i)) && get(i) != math.MaxInt; i = i + i - prev(i) {
			rewards[i] = int(math.Max(float64(rewards[i]), float64(rewards[prev(i)]+1)))
		}
	}

	for _, pit := range pits {
		increase(pit-1, func(a int) int { return a + 1 })
		increase(pit+1, func(a int) int { return a - 1 })
	}

	return rewards
}
