package clonegraph

type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph returns a deep copy of the connected undirected graph.
func cloneGraph(node *Node) *Node {
	return cloneByDFS(node)
}

func cloneByBFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	clones := map[*Node]*Node{
		node: {Val: node.Val, Neighbors: make([]*Node, 0, len(node.Neighbors))},
	}

	queue := []*Node{node}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		currClone := clones[curr]

		for _, nb := range curr.Neighbors {
			if _, ok := clones[nb]; !ok {
				clones[nb] = &Node{Val: nb.Val, Neighbors: make([]*Node, 0, len(nb.Neighbors))}
				queue = append(queue, nb)
			}
			currClone.Neighbors = append(currClone.Neighbors, clones[nb])
		}
	}

	return clones[node]
}

func cloneByDFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	seen := make(map[*Node]*Node, 16)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if clone, ok := seen[node]; ok {
			return clone
		}

		newNode := &Node{
			Val: node.Val,
		}
		seen[node] = newNode

		newNode.Neighbors = make([]*Node, 0, len(node.Neighbors))
		for _, nb := range node.Neighbors {
			newCopy := dfs(nb)
			newNode.Neighbors = append(newNode.Neighbors, newCopy)
		}

		return newNode
	}

	return dfs(node)
}
