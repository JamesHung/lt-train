package rottingoranges

import "fmt"

const debugRottingOranges = false

type Node struct {
	Row    int
	Col    int
	Minute int
}

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// orangesRotting performs multi-source BFS from the initial rotten oranges.
func orangesRotting(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	queue := make([]*Node, 0, row*col)
	freshOrange := 0

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			switch grid[r][c] {
			case 2:
				queue = append(queue, &Node{
					Row:    r,
					Col:    c,
					Minute: 0,
				})
			case 1:
				freshOrange++
			default:
				continue
			}
		}
	}

	if freshOrange == 0 {
		return 0
	}

	minute := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		r, c := node.Row, node.Col
		minute = node.Minute

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= row || nc < 0 || nc >= col {
				continue
			}

			if grid[nr][nc] != 1 {
				continue
			}

			freshOrange--
			grid[nr][nc] = 2
			queue = append(queue, &Node{Row: nr, Col: nc, Minute: node.Minute + 1})
		}
	}

	fmt.Printf("f: %d, m: %d", freshOrange, minute)
	if freshOrange > 0 {
		return -1
	}

	return minute
}
