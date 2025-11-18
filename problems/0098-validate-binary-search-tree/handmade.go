package validatebinarysearchtree

import "math"

const debugValidateBST = false

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST checks whether the tree satisfies strict increasing order constraints.
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, low, high int64) bool {
	if node == nil {
		return true
	}

	value := int64(node.Val)
	if value <= low || value >= high {
		return false
	}

	return validate(node.Left, low, value) && validate(node.Right, value, high)
}
