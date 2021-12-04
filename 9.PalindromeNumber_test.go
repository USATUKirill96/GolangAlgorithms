package main

import (
	"strconv"
	"testing"
)

//Given an integer x, return true if x is palindrome integer.
//An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.

func TestPalindromeNumber(t *testing.T) {

	tests := []struct {
		x      int
		result bool
	}{
		{
			x:      -121,
			result: false,
		},
		{
			x:      121,
			result: true,
		},
		{
			x:      10,
			result: false,
		},
	}

	for _, tt := range tests {
		res := palindromeNumber(tt.x)
		if res != tt.result {
			t.Errorf("Error in 1.TwoSum function. Expected %v, got %v", tt.result, res)
		}
	}
}

func palindromeNumber(x int) bool {
	xstr := strconv.Itoa(x)
	xarr := []rune(xstr)
	for i, j := 0, len(xarr)-1; i < j; i, j = i+1, j-1 {
		xarr[i], xarr[j] = xarr[j], xarr[i]
	}
	return string(xarr) == xstr
}
