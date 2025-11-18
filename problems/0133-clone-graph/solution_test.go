package clonegraph

import "testing"

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name string
		node *Node
	}{
		{
			name: "example",
			node: buildGraph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
		{
			name: "single",
			node: buildGraph([][]int{{}}),
		},
		{
			name: "empty",
			node: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			cloned := cloneGraph(tc.node)
			if !equalGraphs(tc.node, cloned, map[*Node]bool{}, map[*Node]bool{}) {
				t.Fatalf("cloneGraph result differs from original graph structure")
			}
			if tc.node != nil && cloned == tc.node {
				t.Fatalf("expected deep copy, got same pointer")
			}
		})
	}
}

func buildGraph(adj [][]int) *Node {
	if len(adj) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adj))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}
	for i, neighbors := range adj {
		for _, nb := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[nb-1])
		}
	}
	return nodes[0]
}

func equalGraphs(a, b *Node, seenA, seenB map[*Node]bool) bool {
	switch {
	case a == nil && b == nil:
		return true
	case a == nil || b == nil:
		return false
	}

	if seenA[a] && seenB[b] {
		return true
	}

	seenA[a] = true
	seenB[b] = true

	if a.Val != b.Val || len(a.Neighbors) != len(b.Neighbors) {
		return false
	}

	for i := range a.Neighbors {
		if !equalGraphs(a.Neighbors[i], b.Neighbors[i], seenA, seenB) {
			return false
		}
	}

	return true
}
