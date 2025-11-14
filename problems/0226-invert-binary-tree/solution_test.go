package invertbinarytree

import "testing"

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "example1",
			root: tree(
				4,
				tree(2, tree(1, nil, nil), tree(3, nil, nil)),
				tree(7, tree(6, nil, nil), tree(9, nil, nil)),
			),
			want: tree(
				4,
				tree(7, tree(9, nil, nil), tree(6, nil, nil)),
				tree(2, tree(3, nil, nil), tree(1, nil, nil)),
			),
		},
		{
			name: "example2",
			root: tree(2, tree(1, nil, nil), tree(3, nil, nil)),
			want: tree(2, tree(3, nil, nil), tree(1, nil, nil)),
		},
		{
			name: "empty",
			root: nil,
			want: nil,
		},
		{
			name: "unbalanced",
			root: tree(5,
				tree(3, tree(2, nil, nil), nil),
				tree(7, nil, tree(8, nil, nil)),
			),
			want: tree(5,
				tree(7, tree(8, nil, nil), nil),
				tree(3, nil, tree(2, nil, nil)),
			),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := invertTree(copyTree(tc.root))
			if !treesEqual(got, tc.want) {
				t.Fatalf("invertTree() mismatch\n got: %v\nwant: %v", treeToSlice(got), treeToSlice(tc.want))
			}
		})
	}
}

func tree(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

func treesEqual(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val {
		return false
	}
	return treesEqual(a.Left, b.Left) && treesEqual(a.Right, b.Right)
}

func treeToSlice(root *TreeNode) []int {
	out := []int{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			out = append(out, 0)
			return
		}
		out = append(out, node.Val)
		if node.Left != nil || node.Right != nil {
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return out
}
