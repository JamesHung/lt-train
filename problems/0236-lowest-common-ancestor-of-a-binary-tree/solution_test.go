package lowestcommonancestorofabinarytree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		func() struct {
			name string
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
			want *TreeNode
		} {
			root, nodes := sampleTree()
			return struct {
				name string
				root *TreeNode
				p    *TreeNode
				q    *TreeNode
				want *TreeNode
			}{
				name: "example split",
				root: root,
				p:    nodes[5],
				q:    nodes[1],
				want: nodes[3],
			}
		}(),
		func() struct {
			name string
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
			want *TreeNode
		} {
			root, nodes := sampleTree()
			return struct {
				name string
				root *TreeNode
				p    *TreeNode
				q    *TreeNode
				want *TreeNode
			}{
				name: "example ancestor",
				root: root,
				p:    nodes[5],
				q:    nodes[4],
				want: nodes[5],
			}
		}(),
		func() struct {
			name string
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
			want *TreeNode
		} {
			root := &TreeNode{Val: 1}
			root.Left = &TreeNode{Val: 2}
			return struct {
				name string
				root *TreeNode
				p    *TreeNode
				q    *TreeNode
				want *TreeNode
			}{
				name: "two nodes",
				root: root,
				p:    root,
				q:    root.Left,
				want: root,
			}
		}(),
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := lowestCommonAncestor(tc.root, tc.p, tc.q)
			if got != tc.want {
				t.Fatalf("lowestCommonAncestor() = %v, want %v", nodeVal(got), nodeVal(tc.want))
			}
		})
	}
}

func nodeVal(n *TreeNode) any {
	if n == nil {
		return nil
	}
	return n.Val
}

func sampleTree() (*TreeNode, map[int]*TreeNode) {
	nodes := map[int]*TreeNode{}
	makeNode := func(val int) *TreeNode {
		n := &TreeNode{Val: val}
		nodes[val] = n
		return n
	}

	root := makeNode(3)
	root.Left = makeNode(5)
	root.Right = makeNode(1)
	root.Left.Left = makeNode(6)
	root.Left.Right = makeNode(2)
	root.Right.Left = makeNode(0)
	root.Right.Right = makeNode(8)
	root.Left.Right.Left = makeNode(7)
	root.Left.Right.Right = makeNode(4)

	return root, nodes
}
