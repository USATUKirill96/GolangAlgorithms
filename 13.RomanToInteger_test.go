package main

import (
	"testing"
)

//Given an integer x, return true if x is palindrome integer.
//An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.

func TestRomanToInteger(t *testing.T) {

	tests := []struct {
		s      string
		result int
	}{
		{
			s:      "III",
			result: 3,
		},
		{
			s:      "IV",
			result: 4,
		},
		{
			s:      "MCMXCIV",
			result: 1994,
		},
	}

	for _, tt := range tests {
		res := romanToInteger(tt.s)
		if res != tt.result {
			t.Errorf("Error in 1.TwoSum function. Expected %v, got %v", tt.result, res)
		}
	}
}

func romanToInteger(s string) int {
	var mapper = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	arr := []rune(s)
	var iarr []int
	for _, v := range arr {
		iarr = append(iarr, mapper[v])
	}
	summ := 0
	for i, vi := range iarr {
		if i == len(iarr)-1 || vi >= iarr[i+1] {
			summ += vi
		} else {
			summ -= vi
		}
	}
	return summ
}
