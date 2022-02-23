package medium

import (
	"USATUKirill96/leetcode/helpers"
	"testing"
)

func TestBSTTraversal(t *testing.T) {

	tree := &helpers.Node{Key: 10}
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(15)
	tree.Insert(22)

	resInOrder := []int{1, 2, 5, 5, 10, 15, 22}
	resPreOrder := []int{10, 5, 2, 1, 5, 15, 22}
	resPostOrder := []int{1, 2, 5, 5, 22, 15, 10}

	res := BSTTraversalInOrder(tree)
	if !helpers.CompareIntegerLists(res, resInOrder) {
		t.Errorf("Error in BTS TraversalInOrder. Expected: %v, got: %v", resInOrder, res)
	}

	res = BSTTraversalPreOrder(tree)
	if !helpers.CompareIntegerLists(res, resPreOrder) {
		t.Errorf("Error in BTS TraversalPreOrder. Expected: %v, got: %v", resPreOrder, res)
	}

	res = BSTTraversalPostOrder(tree)
	if !helpers.CompareIntegerLists(res, resPostOrder) {
		t.Errorf("Error in BTS TraversalPostOrder. Expected: %v, got: %v", resPostOrder, res)
	}
}

// First add left node, then root, then right
func BSTTraversalInOrder(tree *helpers.Node) []int {
	var values []int

	var traverseReq func(*helpers.Node)
	traverseReq = func(node *helpers.Node) {
		if node.Left != nil {
			traverseReq(node.Left)
		}
		values = append(values, node.Key)

		if node.Right != nil {
			traverseReq(node.Right)
		}
	}

	traverseReq(tree)
	return values
}

// First root, then left, then right
func BSTTraversalPreOrder(tree *helpers.Node) []int {
	var values []int

	var traverseReq func(*helpers.Node)
	traverseReq = func(node *helpers.Node) {
		values = append(values, node.Key)
		if node.Left != nil {
			traverseReq(node.Left)
		}
		if node.Right != nil {
			traverseReq(node.Right)
		}
	}
	traverseReq(tree)
	return values
}

// First left, then right, then root
func BSTTraversalPostOrder(tree *helpers.Node) []int {
	var values []int

	var traverseReq func(*helpers.Node)
	traverseReq = func(node *helpers.Node) {
		if node.Left != nil {
			traverseReq(node.Left)
		}
		if node.Right != nil {
			traverseReq(node.Right)
		}
		values = append(values, node.Key)
	}
	traverseReq(tree)
	return values
}
