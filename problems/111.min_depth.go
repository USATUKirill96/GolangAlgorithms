package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	var minDepthHelper func(root *TreeNode, depth int) int
	minDepthHelper = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth - 1
		}
		return min(minDepthHelper(root.Left, depth+1), minDepthHelper(root.Right, depth+1))
	}
	return minDepthHelper(root, 1)
}
