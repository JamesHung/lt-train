package constructbinarytreefrompreorderandinordertraversal

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []interface{}
	}{
		{
			name:     "example full",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     []interface{}{3, 9, 20, nil, nil, 15, 7},
		},
		{
			name:     "single node",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     []interface{}{-1},
		},
		{
			name:     "skewed left",
			preorder: []int{4, 3, 2, 1},
			inorder:  []int{1, 2, 3, 4},
			want:     []interface{}{4, 3, nil, 2, nil, 1},
		},
		{
			name:     "skewed right",
			preorder: []int{1, 2, 3, 4},
			inorder:  []int{1, 2, 3, 4},
			want:     []interface{}{1, nil, 2, nil, 3, nil, 4},
		},
		{
			name:     "balanced",
			preorder: []int{1, 2, 4, 5, 3, 6, 7},
			inorder:  []int{4, 2, 5, 1, 6, 3, 7},
			want:     []interface{}{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := treeToLevelOrder(buildTree(tc.preorder, tc.inorder))
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("buildTree(%v, %v) levelOrder=%v, want %v", tc.preorder, tc.inorder, got, tc.want)
			}
		})
	}
}

func treeToLevelOrder(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	result := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// Trim trailing nils for cleaner comparison.
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}
