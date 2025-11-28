package binarytreerightsideview

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// debugRightSideView toggles verbose traversal logs.
const debugRightSideView = false

// rightSideView performs a level-order traversal and records the last node of each level.
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			latest := node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if i == levelSize-1 {
				result = append(result, latest)
			}
		}
	}

	return result

}
