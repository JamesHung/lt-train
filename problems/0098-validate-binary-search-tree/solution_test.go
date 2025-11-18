package validatebinarysearchtree

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "example1",
			root: tree(2,
				tree(1, nil, nil),
				tree(3, nil, nil),
			),
			want: true,
		},
		{
			name: "example2",
			root: tree(5,
				tree(1, nil, nil),
				tree(4,
					tree(3, nil, nil),
					tree(6, nil, nil),
				),
			),
			want: false,
		},
		{
			name: "single node",
			root: tree(1, nil, nil),
			want: true,
		},
		{
			name: "invalid because of deep left",
			root: tree(10,
				tree(5,
					tree(2,
						tree(1, nil, nil),
						tree(12, nil, nil),
					),
					tree(7, nil, nil),
				),
				tree(15, nil, tree(20, nil, nil)),
			),
			want: false,
		},
		{
			name: "valid with negatives",
			root: tree(0,
				tree(-3,
					tree(-5, nil, nil),
					tree(-1, nil, nil),
				),
				tree(9,
					tree(5, nil, nil),
					tree(12, nil, nil),
				),
			),
			want: true,
		},
		{
			name: "duplicate violates strict inequality",
			root: tree(2,
				tree(2, nil, nil),
				tree(3, nil, nil),
			),
			want: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := isValidBST(tc.root); got != tc.want {
				t.Fatalf("isValidBST(%v) = %v, want %v", treeToSlice(tc.root), got, tc.want)
			}
		})
	}
}

func tree(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func treeToSlice(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}
	result := []interface{}{}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			result = append(result, nil)
			return
		}
		result = append(result, node.Val)
		if node.Left != nil || node.Right != nil {
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return result
}
