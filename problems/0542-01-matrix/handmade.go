package matrix01

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// updateMatrix returns for each cell the distance to the nearest zero via multi-source BFS.
func updateMatrix(mat [][]int) [][]int {
	rowCount, colCount := len(mat), len(mat[0])
	result := make([][]int, rowCount)
	queue := make([][2]int, 0, rowCount*colCount)

	const maxInt = int((^uint(0)) >> 1)

	for i := 0; i < rowCount; i++ {
		result[i] = make([]int, colCount)
		for j := 0; j < colCount; j++ {

			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				result[i][j] = maxInt
			}
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		r, c := current[0], current[1]
		queue = queue[1:]
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= rowCount || nc < 0 || nc >= colCount {
				continue
			}

			if result[nr][nc] > result[r][c]+1 {
				result[nr][nc] = result[r][c] + 1
				queue = append(queue, [2]int{nr, nc})
			}

		}
	}

	return result

}
