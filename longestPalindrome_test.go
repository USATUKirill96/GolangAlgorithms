package main

import (
	"strings"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	cases := []struct {
		tc  string
		res string
	}{
		{
			"abaxyzzyxb",
			"xyzzyx",
		},
	}

	for _, tt := range cases {
		res := LongestPalindrome(tt.tc)
		if res != tt.res {
			t.Errorf("Input: %v, expected: %v, got: %v", tt.tc, tt.res, res)
		}
	}

}

func comp(f, s int, letters []string) (int, int) {
	if f < 0 || s > len(letters)-1 {
		return f + 1, s - 1
	}
	if letters[f] == letters[s] {
		return comp(f-1, s+1, letters)
	}
	return f + 1, s - 1
}

func LongestPalindrome(input string) string {
	letters := strings.Split(input, "")
	var f, s int
	for i := range letters {
		if i == 0 || i == len(letters)-1 {
			continue
		}
		if letters[i] == letters[i+1] {
			a, b := comp(i, i+1, letters)
			if b-a > s-f {
				f, s = a, b
			}
		}
		if letters[i-1] == letters[i+1] {
			a, b := comp(i-1, i+1, letters)
			if b-a > f-s {
				f, s = a, b
			}
		}
	}

	res := ""
	for ; f <= s; f++ {
		res += letters[f]
	}
	return res
}
