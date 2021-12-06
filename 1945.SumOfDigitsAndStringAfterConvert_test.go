package main

import (
	"fmt"
	"strconv"
	"testing"
)

//You are given a string s consisting of lowercase English letters, and an integer k.

//First, convert s into an integer by replacing each letter with its position in the alphabet
//(i.e., replace 'a' with 1, 'b' with 2, ..., 'z' with 26).
// Then, transform the integer by replacing it with the sum of its digits. Repeat the transform operation k times in total.

//For example, if s = "zbax" and k = 2, then the resulting integer would be 8 by the following operations:

//Convert: "zbax" ➝ "(26)(2)(1)(24)" ➝ "262124" ➝ 262124
//Transform #1: 262124 ➝ 2 + 6 + 2 + 1 + 2 + 4 ➝ 17
//Transform #2: 17 ➝ 1 + 7 ➝ 8

//Return the resulting integer after performing the operations described above.

func TestSumOfDigitsAndString(t *testing.T) {

	tests := []struct {
		s   string
		k   int
		res int
	}{
		{
			s:   "iiii",
			k:   1,
			res: 36,
		},
		{
			s:   "leetcode",
			k:   2,
			res: 6,
		},
		{
			s:   "zbax",
			k:   2,
			res: 8,
		},
		{
			s:   "hvmhoasabaymnmsd",
			k:   1,
			res: 79,
		},
	}

	for _, tt := range tests {
		res := sumOfDigitsAndString(tt.s, tt.k)
		if res != tt.res {
			t.Errorf("Error in 1945. Expected %v, got %v", tt.res, res)
		}
	}
}

func sumOfDigitsAndString(s string, k int) int {

	var ss string
	for _, v := range s {
		ss += fmt.Sprint(int(v) - 96)
	}

	rsum := func(number string) string {
		var intsum int
		for _, val := range number {
			intsum = intsum + int(val-'0')
		}
		return fmt.Sprint(intsum)
	}

	for i := 0; i < k; i++ {
		ss = rsum(ss)
	}

	res, _ := strconv.Atoi(ss)

	return res
}
