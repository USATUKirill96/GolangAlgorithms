package hard

import (
	"sort"
	"testing"
)

func TestSubarraySort(t *testing.T) {
	input := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	res := SubarraySort(input)
	if res[0] != 3 || res[1] != 9 || len(res) != 2 {
		t.Errorf("Expected %v got %v", []int{3, 9}, res)
	}
}

func SubarraySort(input []int) []int {

	var unsorted []int

	for i := range input {
		unsorted = append(unsorted, getUnsorted(i, input)...)
	}

	// Could use self-made algorithm for finding min and max for speed improvement
	sort.Ints(unsorted)
	min, max := unsorted[0], unsorted[len(unsorted)-1]
	var minIdx, maxIdx int

	sort.Ints(input)
	for i, v := range input {
		if v == min {
			minIdx = i
		}
		if v == max {
			maxIdx = i
		}
	}

	return []int{minIdx, maxIdx}

}
func getUnsorted(i int, arr []int) []int {
	var unsorted []int

	if i > 0 {
		if arr[i] < arr[i-1] {
			unsorted = append(unsorted, arr[i], arr[i-1])
		}
	}
	if i < len(arr)-1 {
		if arr[i] > arr[i+1] {
			unsorted = append(unsorted, arr[i], arr[i+1])
		}
	}
	return unsorted
}
