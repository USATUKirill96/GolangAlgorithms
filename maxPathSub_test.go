package main

import (
	"USATUKirill96/leetcode/helpers"
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	tree := new(helpers.BST)
	tree.FromList([]int{4, 5, 2, 1, 3, 6, 7})

	res := MaxPathSum(tree)
	fmt.Println(res)

}

func MaxPathSum(tree *helpers.BST) int {
	_, res := maxPathSumReq(tree.Root)
	return res

}

func maxPathSumReq(node *helpers.Node) (int, int) {
	if node == nil {
		return 0, 0
	}

	// LSB: Max sum as if node was just a branch (with left or right child)
	// LS: max sum if node was a tree (with both children)
	LSB, LS := maxPathSumReq(node.Left)
	RSB, RS := maxPathSumReq(node.Right)

	// Pick the largest branch
	maxChildrenBranch := max(LSB, RSB)
	// If children are negative, we'd better not use them and just use current value as the last node
	maxBranchWithCurrentNode := max(node.Key, maxChildrenBranch+node.Key)

	// What's bigger: just a current node or complete triangle of node and both children?
	maxSumOfBeingTriangle := max(maxBranchWithCurrentNode, LSB+RSB+node.Key)
	// Maybe one of children triangles is bigger than our current triangle itself (like root may be negative)
	maxSumOfTriangleAndChildren := max(maxSumOfBeingTriangle, max(LS, RS))

	return maxBranchWithCurrentNode, maxSumOfTriangleAndChildren
}
