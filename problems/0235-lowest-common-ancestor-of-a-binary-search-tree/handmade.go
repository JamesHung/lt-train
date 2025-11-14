package lowestcommonancestorofabinarysearchtree

// TreeNode represents a BST node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const debugLCABST = false

// lowestCommonAncestor walks the BST to find the split point for p and q.
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	current := root
// 	for current != nil {
// 		if current.Val > p.Val && current.Val > q.Val {
// 			current = current.Left
// 		} else if current.Val < p.Val && current.Val < q.Val {
// 			current = current.Right

// 		} else {
// 			return current
// 		}
// 	}

// 	return current
// }

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
