package main

import (
	"fmt"
	"testing"
)

func TestPermutations(t *testing.T) {
	input := []int{1, 2, 3}
	res := Permutations(input)
	fmt.Println(res)
}

func Permutations(input []int) [][]int {
	var permutations [][]int

	var helper func([]int, []int)
	helper = func(arr, permutation []int) {
		if len(arr) == 0 {
			permutations = append(permutations, permutation)
			return
		}
		for i, v := range arr {
			newArr := make([]int, len(arr))
			copy(newArr, arr)
			newArr[i] = newArr[len(newArr)-1]
			newArr = newArr[:len(newArr)-1]
			newPermutation := append(permutation, v)
			helper(newArr, newPermutation)
		}
	}
	helper(input, []int{})
	return permutations
}
