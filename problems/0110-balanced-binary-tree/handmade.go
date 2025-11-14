package balancedbinarytree

import "math"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	isBalanced := true
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil || !isBalanced {
			return 0
		}

		left := height(node.Left)
		right := height(node.Right)

		if (left-right) > 1 || (right-left) > 1 {
			isBalanced = false
		}

		return 1 + int(math.Max(float64(left), float64(right)))
	}

	height(root)

	return isBalanced
}
