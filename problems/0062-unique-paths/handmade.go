package uniquepaths

const debugUniquePaths = false

// uniquePaths returns the number of paths that move only right/down on an m by n grid.
// Uses 1D DP: row[c] accumulates ways to reach current row, column c.
func uniquePaths(m int, n int) int {
	row := make([]int, n)
	for c := range row {
		row[c] = 1
	}

	for i := 1; i < m; i++ {
		for c := 1; c < n; c++ {
			row[c] += row[c-1]
		}
	}

	return row[n-1]
}
