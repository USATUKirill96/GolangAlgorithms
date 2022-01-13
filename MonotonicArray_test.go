package main

import (
	"testing"
)

func TestMonotonicArray(t *testing.T) {

	tests := []struct {
		input  []int
		result bool
	}{
		{
			input:  []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001},
			result: true,
		},
		{
			input:  []int{-1, -5, -3},
			result: false,
		},
	}

	for _, tt := range tests {
		res := MonotonicArray(tt.input)
		if res != tt.result {
			t.Errorf("Error in MonotonicArray function. Expected %v, got %v", tt.result, res)
		}
	}
}

func MonotonicArray(arr []int) bool {
	prev := arr[0]
	// 0 : not set
	// 1 : inc
	// 2: dec
	mode := 0
	for _, v := range arr {
		switch mode {
		case 0:
			if prev > v {
				mode = 2
				prev = v
			} else if prev < v {
				mode = 1
				prev = v
			} else {
				prev = v
			}
		case 1:
			if v >= prev {
				continue
			} else {
				return false
			}
		case 2:
			if v <= prev {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
