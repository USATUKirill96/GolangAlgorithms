package helpers

import (
	"fmt"
	"testing"
)

func TestSuffixTreeInsert(t *testing.T) {
	tree := &SuffixTree{}
	tree.Insert("babc")
	fmt.Print(tree)
}

func TestSuffixTreeSearch(t *testing.T) {
	tree := &SuffixTree{}
	tree.Insert("babc")

	testCases := []struct {
		word   string
		result bool
	}{
		{
			"abc",
			true,
		},
		{
			"c",
			true,
		},
		{
			"bc",
			true,
		},
		{
			"bab",
			false,
		},
	}

	for _, tcase := range testCases {
		res := tree.Search(tcase.word)
		if tcase.result != res {
			t.Errorf("substring: %v expected: %v, got: %v", tcase.word, tcase.result, res)
		}
	}
}
