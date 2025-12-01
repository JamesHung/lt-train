package constructbinarytreefrompreorderandinordertraversal

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const debugBuildTree = false

// buildTree reconstructs the tree from preorder and inorder traversals.
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	index := make(map[int]int, len(inorder))
	for i, v := range inorder {
		index[v] = i
	}

	rootIdx := 0
	var buildFunc func(left int, right int) *TreeNode
	buildFunc = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		value := preorder[rootIdx]
		rootIdx++
		node := &TreeNode{Val: value}
		mid := index[value]

		node.Left = buildFunc(left, mid-1)
		node.Right = buildFunc(mid+1, right)

		return node
	}

	return buildFunc(0, len(preorder)-1)

}
