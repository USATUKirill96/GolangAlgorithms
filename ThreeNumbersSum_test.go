package main

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"
)

func TestThreeNumbersSum(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		result [][]int
	}{
		{
			nums:   []int{12, 3, 1, 2, -6, 5, -8, 6},
			target: 0,
			result: [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}},
		},
	}

	for _, tt := range tests {
		res := ThreeNumbersSumAsync(tt.nums, tt.target)
		t.Logf(fmt.Sprint(res))
	}
}

//func ThreeNumbersSum(arr []int, t int) [][]int {
//sort.Ints(arr[:])

//res := [][]int{}

//var CN, L, R, sum int

//for CN < len(arr) {
//L = CN + 1
//R = len(arr) - 1
//for L < R {
//sum = arr[CN] + arr[L] + arr[R]
//if sum == t {
//res = append(res, []int{arr[CN], arr[L], arr[R]})
//L += 1
//R -= 1
//} else if sum > t {
//R -= 1
//} else {
//L += 1
//}
//}
//CN += 1
//}
//return res
//}

type threeNumbersSumResult struct {
	mu                 sync.Mutex
	returnValue        [][]int
	goroutinesFinished int
}

func (t *threeNumbersSumResult) add(res []int) {
	t.mu.Lock()
	t.returnValue = append(t.returnValue, res)
	t.mu.Unlock()
}

func (t *threeNumbersSumResult) finish() {
	t.mu.Lock()
	t.goroutinesFinished += 1
	t.mu.Unlock()
}

func (t *threeNumbersSumResult) isFinished(total int) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.goroutinesFinished == total
}

func ThreeNumbersSumAsync(arr []int, t int) [][]int {
	sort.Ints(arr[:])

	res := &threeNumbersSumResult{}

	getTriplets := func(CN int) {
		var L, R, sum int
		L = CN + 1
		R = len(arr) - 1
		for L < R {
			sum = arr[CN] + arr[L] + arr[R]
			if sum == t {
				fmt.Println(sum)
				res.add([]int{arr[CN], arr[L], arr[R]})
				L += 1
				R -= 1
			} else if sum > t {
				R -= 1
			} else {
				L += 1
			}
		}
		res.finish()
	}

	var CN int
	for CN < len(arr) {
		go getTriplets(CN)
		CN += 1
	}

	for !res.isFinished(len(arr)) {
		time.Sleep(time.Microsecond)
	}

	return res.returnValue
}
