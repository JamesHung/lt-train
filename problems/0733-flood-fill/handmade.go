package floodfill

const debugFloodFill = false

var directions = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// floodFill repaints all connected pixels that match the starting color.
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	original := image[sr][sc]
	if original == color {
		return image
	}

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
			return
		}
		if image[r][c] != original {
			return
		}

		image[r][c] = color

		for _, dir := range directions {
			dfs(r+dir[0], c+dir[1])
		}
	}

	dfs(sr, sc)

	return image
}
