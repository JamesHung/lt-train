package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const useBFS = false

func levelOrder(root *TreeNode) [][]int {
	if useBFS {
		return bfsLevelOrder(root)
	}

	return dfsLevelOrder(root)
}

// levelOrder returns a slice of levels filled with node values in BFS order.
func bfsLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	// queue := make([]*TreeNode, 0, 1)
	// queue = append(queue, root)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelVals = append(levelVals, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelVals)
	}

	return result
}

func dfsLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(result) == level {
			result = append(result, []int{})
		}

		result[level] = append(result[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return result
}
