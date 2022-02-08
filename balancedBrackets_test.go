package main

import (
	"strings"
	"testing"
)

var backetsTyping = map[string]string{
	"(": "opening",
	"[": "opening",
	"{": "opening",
	")": "closing",
	"]": "closing",
	"}": "closing",
}
var bracketsMatching = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	")": "(",
	"]": "[",
	"}": "{",
}

func TestBalancedBrackets(t *testing.T) {

	cases := []struct {
		tc  string
		res bool
	}{
		{
			"(([]{}[{}]){})",
			true,
		},
		{
			"(]",
			false,
		},
		{
			tc:  "(",
			res: false,
		},
		{
			tc:  ")",
			res: false,
		},
		{
			"",
			true,
		},
	}

	for _, tt := range cases {
		res := BalancedBrackets(tt.tc)
		if res != tt.res {
			t.Errorf("Input: %v, expected: %v, got: %v", tt.tc, tt.res, res)
		}
	}

}

func BalancedBrackets(input string) bool {
	stack := &bracketsStack{}
	for _, v := range strings.Split(input, "") {
		if backetsTyping[v] == "opening" {
			stack.push(v)
		} else {
			if stack.isEmpty() {
				return false
			}
			opening := stack.pop()
			if !corresponds(opening, v) {
				return false
			}
		}
	}
	return stack.isEmpty()
}

type bracketsStack struct {
	brackets []string
}

func (st *bracketsStack) pop() string {
	value := st.brackets[(len(st.brackets))-1]
	st.brackets = st.brackets[:len(st.brackets)-1]
	return value
}

func (st *bracketsStack) push(val string) {
	st.brackets = append(st.brackets, val)
}

func (st *bracketsStack) isEmpty() bool {
	return len(st.brackets) == 0
}

func corresponds(opening, closing string) bool {
	return bracketsMatching[opening] == closing
}
