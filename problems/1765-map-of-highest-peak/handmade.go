package mapofhighestpeak

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// highestPeak assigns heights via multi-source BFS starting from all water cells.
func highestPeak(isWater [][]int) [][]int {
	rowCnt, colCnt := len(isWater), len(isWater[0])
	result := make([][]int, rowCnt)
	queue := make([][2]int, 0, rowCnt*colCnt)

	for r := 0; r < rowCnt; r++ {
		result[r] = make([]int, colCnt)
		for c := 0; c < colCnt; c++ {
			if isWater[r][c] == 1 {
				result[r][c] = 0
				queue = append(queue, [2]int{r, c})
			} else {
				result[r][c] = -1
			}
		}
	}

	head := 0
	for head < len(queue) {
		node := queue[head]
		head++
		nodeR, nodeC := node[0], node[1]
		for _, dir := range directions {
			nearR, nearC := nodeR+dir[0], nodeC+dir[1]
			if nearR < 0 || nearR >= rowCnt || nearC < 0 || nearC >= colCnt {
				continue
			}

			if result[nearR][nearC] == -1 {
				result[nearR][nearC] = result[nodeR][nodeC] + 1
				queue = append(queue, [2]int{nearR, nearC})
			}
		}
	}

	return result
}
