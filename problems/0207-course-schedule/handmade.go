package courseschedule

const debugCanFinish = false

// canFinish uses Kahn's algorithm (BFS topological sort) to detect cycles.
//
//	numCourses = 2, prerequisites = [[1,0]]
func canFinish(numCourses int, prerequisites [][]int) bool {

	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	queue := make([]int, 0, numCourses)

	for _, edge := range prerequisites {
		inDegreeNode := edge[0]
		outDegreeNode := edge[1]
		graph[outDegreeNode] = append(graph[outDegreeNode], inDegreeNode)
		inDegrees[inDegreeNode]++
	}

	for course, degree := range inDegrees {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	visited := 0
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		visited++
		for _, next := range graph[current] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return visited == numCourses
}
