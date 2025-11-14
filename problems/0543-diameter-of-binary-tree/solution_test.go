package diameterofbinarytree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example1",
			root: tree(1,
				tree(2,
					tree(4, nil, nil),
					tree(5, nil, nil),
				),
				tree(3, nil, nil),
			),
			want: 3,
		},
		{
			name: "example2",
			root: tree(1, tree(2, nil, nil), nil),
			want: 1,
		},
		{
			name: "single node",
			root: tree(42, nil, nil),
			want: 0,
		},
		{
			name: "skewed tree",
			root: tree(1,
				tree(2,
					tree(3,
						tree(4, nil, nil),
						nil,
					),
					nil,
				),
				nil,
			),
			want: 3,
		},
		{
			name: "diameter through mixed path",
			root: tree(1,
				tree(2,
					tree(4,
						tree(7, nil, nil),
						nil,
					),
					tree(5, nil, nil),
				),
				tree(3,
					nil,
					tree(6, nil, nil),
				),
			),
			want: 5,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := diameterOfBinaryTree(tc.root)
			if got != tc.want {
				t.Fatalf("diameterOfBinaryTree(%v) = %d, want %d", treeToSlice(tc.root), got, tc.want)
			}
		})
	}
}

func tree(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func treeToSlice(root *TreeNode) []interface{} {
	out := []interface{}{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			out = append(out, nil)
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
