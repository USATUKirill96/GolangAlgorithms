package main

import (
	"fmt"
	"testing"
)

func TestFourNumberSum(t *testing.T) {
	input := []int{7, 6, 4, -1, 1, 2}
	target := 16
	res := fourNumberSum(input, target)
	fmt.Println(res)
}

func fourNumberSum(input []int, target int) [][]int {
	var result [][]int
	complements := make(map[int][][]int)

	for numberIdx := range input {
		// split through all the numbers after current
		for postNumberIdx := numberIdx + 1; postNumberIdx < len(input); postNumberIdx++ {
			// current pair is []int{input[numberIdx], input[postNumberIdx]}
			complement, ok := complements[target-(input[numberIdx]+input[postNumberIdx])]
			if ok {
				// Add all possible complements to current pair and append to result
				for _, pair := range complement {
					result = append(result,
						append(pair, input[numberIdx], input[postNumberIdx]))
				}
			}
		}

		// split through all the numbers before current
		for prevNumberIdx := numberIdx - 1; prevNumberIdx >= 0; prevNumberIdx-- {
			// creating pairs of numbers using current number and each of the previous
			complementTarget := input[numberIdx] + input[prevNumberIdx]
			complementPair := []int{input[numberIdx], input[prevNumberIdx]}

			// In the hashmap of complements check if there is already a pair making this sum
			existingTarget, ok := complements[complementTarget]
			if ok {
				// if yes, just add this pair too
				existingTarget = append(existingTarget, complementPair)
			} else {
				// else create it
				complements[complementTarget] = [][]int{complementPair}
			}

		}

	}
	return result
}
