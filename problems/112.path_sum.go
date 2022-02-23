package problems

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func hasPathSum(root *treeNode, targetSum int) bool {
	found := false
	hasPathSumHelper(root, 0, targetSum, &found)
	return found
}

func hasPathSumHelper(root *treeNode, prevSum, required int, found *bool) {
	if root == nil {
		return
	}
	if root.Left != nil {
		hasPathSumHelper(root.Left, prevSum+root.Val, required, found)
	}
	if root.Right != nil {
		hasPathSumHelper(root.Right, prevSum+root.Val, required, found)
	}
	if root.Left == nil && root.Right == nil {
		if prevSum+root.Val == required {
			*found = true
		}
	}

}
