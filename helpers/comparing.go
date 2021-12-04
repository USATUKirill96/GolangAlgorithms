package helpers

func CompareIntegerLists(a []int, b []int) bool {
	if len(a) == 0 || len(b) == 0 {
		if len(a) == len(b) {
			return true
		}
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
