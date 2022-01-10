package main

//type TreeNode struct {
//Val   int
//Left  *TreeNode
//Right *TreeNode
//}

func maxDepth(root *TreeNode) int {

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	var maxDepthHelper func(root *TreeNode, depth int) int
	maxDepthHelper = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth - 1
		}
		return max(maxDepthHelper(root.Left, depth+1), maxDepthHelper(root.Right, depth+1))
	}
	return maxDepthHelper(root, 1)
}
