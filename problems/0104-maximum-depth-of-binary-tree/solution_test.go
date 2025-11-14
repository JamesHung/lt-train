package maximumdepthofbinarytree

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "example1",
			root: tree(3,
				tree(9, nil, nil),
				tree(20, tree(15, nil, nil), tree(7, nil, nil)),
			),
			want: 3,
		},
		{
			name: "example2",
			root: tree(1, nil, tree(2, nil, nil)),
			want: 2,
		},
		{
			name: "empty tree",
			root: nil,
			want: 0,
		},
		{
			name: "skewed left",
			root: tree(5,
				tree(4,
					tree(3,
						tree(2, nil, nil),
						nil,
					),
					nil,
				),
				nil,
			),
			want: 4,
		},
		{
			name: "balanced deeper",
			root: tree(1,
				tree(2,
					tree(4, nil, nil),
					tree(5, nil, nil),
				),
				tree(3,
					tree(6, nil, nil),
					tree(7, nil, tree(8, nil, nil)),
				),
			),
			want: 4,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := maxDepth(tc.root); got != tc.want {
				t.Fatalf("maxDepth(%v) = %d, want %d", treeToSlice(tc.root), got, tc.want)
			}
		})
	}
}

func tree(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func treeToSlice(root *TreeNode) []interface{} {
	out := []interface{}{}
	var dfs func(*TreeNode)
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
