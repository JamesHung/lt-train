package balancedbinarytree

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			name: "example balanced",
			root: tree(3,
				tree(9, nil, nil),
				tree(20, tree(15, nil, nil), tree(7, nil, nil)),
			),
			want: true,
		},
		{
			name: "example unbalanced",
			root: tree(1,
				tree(2,
					tree(3,
						tree(4, nil, nil),
						tree(4, nil, nil),
					),
					tree(3, nil, nil),
				),
				tree(2, nil, nil),
			),
			want: false,
		},
		{
			name: "empty",
			root: nil,
			want: true,
		},
		{
			name: "single node",
			root: tree(42, nil, nil),
			want: true,
		},
		{
			name: "skewed",
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
			want: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := isBalanced(tc.root)
			if got != tc.want {
				t.Fatalf("isBalanced() = %v, want %v", got, tc.want)
			}
		})
	}
}

func tree(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}
