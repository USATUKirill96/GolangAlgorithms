package medium

import (
	"USATUKirill96/leetcode/helpers"
	"fmt"
	"testing"
)

func TestRemoveNthNodeFromEnd(t *testing.T) {
	input := []int{0, 1, 2}
	linkedList := &helpers.LinkedList{}
	linkedList.FromList(input)
	fmt.Println(linkedList)
	RemoveNthNodeFromEnd(linkedList, 4)
	fmt.Println(linkedList)
}

func RemoveNthNodeFromEnd(input *helpers.LinkedList, n int) {

	prev := input.Root
	first := input.Root
	second := input.Root

	for i := 0; i < n; i++ {
		if second == nil {
			break
		}
		second = second.Next
	}

	for second != nil {
		prev = first
		first = first.Next
		second = second.Next

	}

	if prev == first {
		input.Root = first.Next
		return
	}
	prev.Next = first.Next

}
