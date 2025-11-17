package matrix01

const debugUpdateMatrix = false

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// updateMatrix returns for each cell the distance to the nearest zero via multi-source BFS.
func updateMatrix(mat [][]int) [][]int {
	rowLen, colLen := len(mat), len(mat[0])
	result := make([][]int, rowLen)
	queue := make([][2]int, 0, rowLen*colLen)

	const max_int = int(^uint(0) >> 1)

	// Find all zero cells to use as BFS sources.
	for row := 0; row < rowLen; row++ {
		result[row] = make([]int, colLen)
		for col := 0; col < colLen; col++ {
			if mat[row][col] == 0 {
				result[row][col] = 0
				queue = append(queue, [2]int{row, col})
			} else {
				result[row][col] = max_int
			}
		}
	}

	// BFS from all zero cells
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodeRow, nodeCol := node[0], node[1]

		for _, direction := range directions {
			nearR, nearC := nodeRow+direction[0], nodeCol+direction[1]
			if nearR < 0 || nearR >= rowLen || nearC < 0 || nearC >= colLen {
				continue
			}

			if result[nearR][nearC] > result[nodeRow][nodeCol]+1 {
				result[nearR][nearC] = result[nodeRow][nodeCol] + 1
				queue = append(queue, [2]int{nearR, nearC})
			}
		}

	}

	return result

}
