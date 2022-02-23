package hard

import "testing"

func TestLargestRanges(t *testing.T) {
	input := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
	res := LargestRange(input)
	expected := []int{0, 7}
	if res[0] != expected[0] || res[1] != expected[1] || len(res) != 2 {
		t.Errorf("expected %v, got %v", expected, res)
	}
}

func LargestRange(input []int) []int {

	exploredNumbers := make(map[int]bool)
	for _, v := range input {
		exploredNumbers[v] = false
	}

	largestRange := []int{0, 0}

	for _, v := range input {
		// if current number was already explored, just skip it
		if exploredNumbers[v] {
			continue
		}

		// find left and right borders of range which includes the current number
		leftBorder := getBorder(exploredNumbers, v, func(a int) int { return a - 1 })
		rightBorder := getBorder(exploredNumbers, v, func(a int) int { return a + 1 })
		exploredNumbers[v] = true

		// if new range is greater than the previous, replace it
		if rightBorder-leftBorder > largestRange[1]-largestRange[0] {
			largestRange[0], largestRange[1] = leftBorder, rightBorder
		}

	}
	return largestRange
}

func getBorder(exploredNumbers map[int]bool, position int, iterate func(int) int) int {
	// increase or decrease previous number
	currentPosition := iterate(position)

	// if this number is not in the input, return the previous number which was
	_, exists := exploredNumbers[currentPosition]
	if !exists {
		return position
	}

	// if this number is in the input, mark it as visited, so not to split over it in the future, and try next
	exploredNumbers[currentPosition] = false
	return getBorder(exploredNumbers, currentPosition, iterate)
}
