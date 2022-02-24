package hard

import (
	"USATUKirill96/leetcode/helpers"
	"testing"
)

func TestWaterArea(t *testing.T) {
	input := []int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3}
	expected := 48
	res := WaterArea(input)
	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}

func WaterArea(levels []int) int {
	leftMax := make([]int, len(levels))
	rightMax := make([]int, len(levels))
	leftMax[0] = 0
	rightMax[len(rightMax)-1] = 0

	for i := 1; i < len(levels); i++ {
		leftMax[i] = helpers.Max(leftMax[i-1], levels[i-1])
		rightIdx := len(levels) - 1 - i
		rightMax[rightIdx] = helpers.Max(rightMax[rightIdx+1], levels[rightIdx+1])
	}

	var totalAmountOfWater int
	for i := range levels {
		waterLevel := helpers.Min(leftMax[i], rightMax[i])
		if waterLevel > levels[i] {
			totalAmountOfWater += waterLevel - levels[i]
		}
	}
	return totalAmountOfWater
}
