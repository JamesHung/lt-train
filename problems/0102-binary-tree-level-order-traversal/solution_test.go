package binarytreelevelordertraversal

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "example1",
			root: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "single",
			root: &TreeNode{Val: 1},
			want: [][]int{{1}},
		},
		{
			name: "empty",
			root: nil,
			want: nil,
		},
		{
			name: "skewed",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			want: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := levelOrder(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("levelOrder(...) = %v, want %v", got, tc.want)
			}
		})
	}
}
