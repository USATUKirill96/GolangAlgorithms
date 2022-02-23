package hard

import (
	"strings"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	str1 := "zxvvyzw"
	str2 := "xkykzpw"
	expected := "xyzw"
	result := longestCommonSubsequence(str1, str2)
	if expected != result {
		t.Errorf("expected %v, got %v", expected, result)
	}

}

func longestCommonSubsequence(str1, str2 string) string {

	// Splitting input into list of headers
	columns := strings.Split(str1, "")
	rows := strings.Split(str2, "")

	// Prepending empty string before headers for edge cases
	columns = append([]string{""}, columns...)
	rows = append([]string{""}, rows...)

	//    | "" | A | B | C |
	// "" |
	// A  |
	// C  |
	// Creating the table of solution
	table := make([][]string, len(rows))
	for rowIdx := range table {
		table[rowIdx] = make([]string, len(columns))
	}

	//    | "" | A | B | C |
	// "" | "" |"" | ""| ""|
	// A  | "" |
	// C  | "" |
	// Filling the edge cases
	for row := range table {
		table[row][0] = ""
	}
	for column := range table[0] {
		table[0][column] = ""
	}

	for row := 1; row < len(table); row++ {
		for column := 1; column < len(table[row]); column++ {
			biggestNeighbor := getBiggestNeighbor(table, row, column)
			// The main logic of the algorithm
			if rows[row] == columns[column] {
				table[row][column] = biggestNeighbor + rows[row]
			} else {
				table[row][column] = biggestNeighbor
			}
		}
	}

	return table[len(table)-1][len(table[0])-1]
}

//    | "" | X | Z | Y |
// "" | "" |"" | ""| ""|
// X  | "" | X | X if len(X) > len("") else "" |
// Y  | "" |
// Z  | "" |
func getBiggestNeighbor(table [][]string, row, column int) string {
	upper := table[row-1][column]
	left := table[row][column-1]
	if len(upper) > len(left) {
		return upper
	} else {
		return left
	}
}
