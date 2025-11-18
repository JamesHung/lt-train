package numberofislands

const useDFS = false

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// numIslands counts connected land components using DFS.
func numIslands(grid [][]byte) int {
	if useDFS {
		return checkByDFS(grid)
	}

	return checkByBFS(grid)
}

func checkByBFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	bfs := func(r int, c int) {
		if grid[r][c] != '1' {
			return
		}
		count++
		grid[r][c] = '0'
		queue := [][2]int{{r, c}}

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			r, c := node[0], node[1]

			for _, dir := range directions {
				nr := r + dir[0]
				nc := c + dir[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				nrValue := grid[nr][nc]
				if nrValue == '1' {
					queue = append(queue, [2]int{nr, nc})
					grid[nr][nc] = '0'
				}
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			bfs(r, c)
		}
	}

	return count

}

func checkByDFS(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	var dfs func(r int, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}

		grid[r][c] = '0'
		for _, dir := range directions {
			dfs(r+dir[0], c+dir[1])
		}

	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}
