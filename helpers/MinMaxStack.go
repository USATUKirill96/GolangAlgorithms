package helpers

import (
	"errors"
	"math"
)

type MinMaxStack struct {
	peeks []int
	min   []int
	max   []int
}

func (st *MinMaxStack) Max() int {
	if len(st.max) == 0 {
		return math.MinInt
	}
	return st.max[len(st.max)-1]
}

func (st *MinMaxStack) Min() int {
	if len(st.min) == 0 {
		return math.MaxInt
	}
	return st.min[len(st.min)-1]
}

func (st *MinMaxStack) Push(i int) {

	st.peeks = append(st.peeks, i)

	min, _ := order(st.Min(), i)
	st.min = append(st.min, min)

	_, max := order(st.Max(), i)
	st.max = append(st.max, max)

}

func (st *MinMaxStack) Pop() (int, error) {
	if len(st.peeks) == 0 {
		return 0, errors.New("Stack is empty")
	}

	value := st.peeks[len(st.peeks)-1]

	st.peeks = st.peeks[:len(st.peeks)-1]
	st.min = st.min[:len(st.min)-1]
	st.max = st.max[:len(st.max)-1]

	return value, nil
}

func order(a, b int) (int, int) {
	if a <= b {
		return a, b
	} else {
		return b, a
	}
}
