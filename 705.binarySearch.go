package main

func search(nums []int, target int) int {
	var m int
	var pv int
	l := 0
	r := len(nums) - 1

	for l <= r {
		m = (l + r) / 2
		pv = nums[m]
		if pv == target {
			return m
		} else if pv < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
