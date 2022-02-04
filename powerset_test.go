package main

import (
	"fmt"
	"testing"
)

func TestPowerset(t *testing.T) {
	input := []int{1, 2, 3}
	res := Powerset(input)
	fmt.Println(res)

}

func Powerset(input []int) [][]int {
	result := [][]int{{}}
	for _, v := range input {
		var newSets [][]int
		for _, s := range result {
			newSet := append(s, v)
			newSets = append(newSets, newSet)
		}
		result = append(result, newSets...)
	}
	return result
}
