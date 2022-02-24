package hard

import (
	"USATUKirill96/leetcode/helpers"
	"math"
	"testing"
)

func TestMinNumberOfJumps(t *testing.T) {
	input := []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3}
	expected := 4
	result := MinNumberOfJumps(input)
	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func MinNumberOfJumps(arr []int) int {

	// jumps[i] = number of jumps to get from arr[0] to arr[i]
	var jumps []int
	for range arr {
		// init with MaxInt values for future needs
		jumps = append(jumps, math.MaxInt-1)
	}
	// We need no jumps to get from arr[0] to arr[0]
	jumps[0] = 0

	for i := range arr {
		for j := i; j >= 0; j-- {
			// if you can reach arr[i] from arr [j]
			if arr[j]+j >= i {
				// Even though you can reach arr[i] from arr[j] it's not sure the new way is shorter than, for example, from arr[j+1]
				jumps[i] = helpers.Min(jumps[i], jumps[j]+1)
			}
		}
	}

	return jumps[len(jumps)-1]
}
