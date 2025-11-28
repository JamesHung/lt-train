package binarytreerightsideview

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{
			name: "example1",
			root: n(1,
				n(2, nil, n(5, nil, nil)),
				n(3, nil, n(4, nil, nil)),
			),
			want: []int{1, 3, 4},
		},
		{
			name: "example2",
			root: n(1,
				n(2,
					n(4, n(5, nil, nil), nil),
					nil,
				),
				n(3, nil, nil),
			),
			want: []int{1, 3, 4, 5},
		},
		{
			name: "example3",
			root: n(1, nil, n(3, nil, nil)),
			want: []int{1, 3},
		},
		{
			name: "empty tree",
			root: nil,
			want: nil,
		},
		{
			name: "left heavy",
			root: n(1,
				n(2, n(3, n(4, nil, nil), nil), nil),
				nil,
			),
			want: []int{1, 2, 3, 4},
		},
		{
			name: "zigzag deeper right",
			root: n(1,
				n(2, n(3, nil, n(6, nil, nil)), n(5, nil, nil)),
				n(4, nil, n(7, nil, nil)),
			),
			want: []int{1, 4, 7, 6},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := rightSideView(tc.root); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("rightSideView() = %v, want %v", got, tc.want)
			}
		})
	}
}

func n(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}
