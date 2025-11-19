package binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const useBFS = true

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
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			levelVals = append(levelVals, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
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
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(result) == level {
			result = append(result, []int{})
		}

		result[level] = append(result[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return result
}
