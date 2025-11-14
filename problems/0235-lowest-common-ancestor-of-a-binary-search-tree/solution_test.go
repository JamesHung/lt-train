package lowestcommonancestorofabinarysearchtree

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
			root, nodes := sampleBST()
			return struct {
				name string
				root *TreeNode
				p    *TreeNode
				q    *TreeNode
				want *TreeNode
			}{
				name: "example split",
				root: root,
				p:    nodes[2],
				q:    nodes[8],
				want: nodes[6],
			}
		}(),
		func() struct {
			name string
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
			want *TreeNode
		} {
			root, nodes := sampleBST()
			return struct {
				name string
				root *TreeNode
				p    *TreeNode
				q    *TreeNode
				want *TreeNode
			}{
				name: "example ancestor is node",
				root: root,
				p:    nodes[2],
				q:    nodes[4],
				want: nodes[2],
			}
		}(),
		func() struct {
			name string
			root *TreeNode
			p    *TreeNode
			q    *TreeNode
			want *TreeNode
		} {
			root := &TreeNode{Val: 2}
			root.Left = &TreeNode{Val: 1}
			return struct {
				name string
				root *TreeNode
				p    *TreeNode
				q    *TreeNode
				want *TreeNode
			}{
				name: "two node tree",
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

func sampleBST() (*TreeNode, map[int]*TreeNode) {
	nodes := map[int]*TreeNode{}
	makeNode := func(val int) *TreeNode {
		n := &TreeNode{Val: val}
		nodes[val] = n
		return n
	}

	root := makeNode(6)
	root.Left = makeNode(2)
	root.Right = makeNode(8)
	root.Left.Left = makeNode(0)
	root.Left.Right = makeNode(4)
	root.Right.Left = makeNode(7)
	root.Right.Right = makeNode(9)
	root.Left.Right.Left = makeNode(3)
	root.Left.Right.Right = makeNode(5)

	return root, nodes
}

func nodeVal(n *TreeNode) any {
	if n == nil {
		return nil
	}
	return n.Val
}
