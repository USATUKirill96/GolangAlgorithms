package problems

import (
	"testing"
)

func Test982(t *testing.T) {

	tests := []struct {
		s string
		t string
		r bool
	}{
		{
			"abc",
			"ahbgdc",
			true,
		},
		{
			"axc",
			"ahbgdc",
			false,
		},
		{
			"",
			"aaabbc",
			true,
		},
		{
			"a",
			"b",
			false,
		},
	}

	for _, tt := range tests {
		res := isSubsequence(tt.s, tt.t)
		if res != tt.r {
			t.Errorf("Error in 982. Expected %v, got %v", tt.r, res)
		}
	}
}

func isSubsequence(s, t string) bool {
	var sIdx, tIdx int

	sarr := []rune(s)
	tarr := []rune(t)

	for sIdx < len(sarr) {
		if tIdx == len(tarr) {
			return false
		}
		if tarr[tIdx] == sarr[sIdx] {
			sIdx += 1
		}
		tIdx += 1
	}
	return true
}
