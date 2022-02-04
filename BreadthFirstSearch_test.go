package main

import (
	"USATUKirill96/leetcode/helpers"
	"fmt"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	tt := helpers.BST{}
	tt.FromList([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(tt)
	res := BreadthFirstSearch(tt)
	fmt.Println(res)
	if !helpers.CompareIntegerLists(res, []int{4, 2, 6, 1, 3, 5, 7}) {
		t.Error("Error in TestBreadthFirstSearch")
	}

}

type queue struct {
	nodes []*helpers.Node
}

func (q *queue) add(nodes ...*helpers.Node) {
	for _, n := range nodes {
		if n != nil {
			q.nodes = append(q.nodes, n)
		}
	}
}
func (q *queue) pop() *helpers.Node {
	val := q.nodes[0]
	q.nodes = q.nodes[1:]
	return val
}
func (q *queue) isEmpty() bool {
	return len(q.nodes) == 0
}

func BreadthFirstSearch(t helpers.BST) []int {
	var result []int
	var currentNode *helpers.Node
	q := &queue{}
	q.add(t.Root)
	for !q.isEmpty() {
		currentNode = q.pop()
		result = append(result, currentNode.Key)
		q.add(currentNode.Left, currentNode.Right)
	}
	return result
}
