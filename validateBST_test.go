package main

import (
	"USATUKirill96/leetcode/helpers"
	"math"
	"testing"
)

func TestValidateBST(t *testing.T) {
	treeValues := []int{5, 15, 2, 5, 1, 13, 22, 14}
	tree := &helpers.Node{Key: 10}
	for _, v := range treeValues {
		tree.Insert(v)
	}

	if ValidateBST(tree) != true {
		t.Error("Invalid result in validate tree")
	}

	// mess up the tree
	tree.Get(13).Key = 16

	if ValidateBST(tree) != false {
		t.Error("Invalid result in validate tree")
	}
}

func ValidateBST(tree *helpers.Node) bool {
	result := true

	var validate func(*helpers.Node, int, int)
	validate = func(node *helpers.Node, min, max int) {
		if node.Key < min || node.Key >= max {
			result = false
		}
		if node.Right != nil {
			validate(node.Right, node.Key, max)
		}
		if node.Left != nil {
			validate(node.Left, min, node.Key)
		}
	}
	validate(tree, -math.MaxInt, math.MaxInt)
	return result
}
