package medium

import (
	"fmt"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	input := []string{"yo", "act", "blop", "tac", "cat", "oy", "olbp"}
	res := GroupAnagrams(input)
	fmt.Println(res)
}

func GroupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)

	for _, word := range words {
		runes := SortRuneString(word)
		sort.Sort(runes)
		key := string(runes)

		_, exists := groups[key]
		if exists {
			groups[key] = append(groups[key], word)
		} else {
			groups[key] = []string{word}
		}
	}

	var result [][]string
	for _, v := range groups {
		result = append(result, v)
	}

	return result
}

type SortRuneString []rune

func (s SortRuneString) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortRuneString) Less(i, j int) bool { return s[i] < s[j] }
func (s SortRuneString) Len() int           { return len(s) }
