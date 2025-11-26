package spiralmatrix

import "fmt"

const debugSpiralMatrix = false

// spiralOrder returns all elements of the matrix in spiral order.
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		top++
		if top > bottom {
			break
		}

		for r := top; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		right--
		if left > right {
			break
		}

		for c := right; c >= left; c-- {
			res = append(res, matrix[bottom][c])
		}
		bottom--
		if top > bottom {
			break
		}

		for r := bottom; r >= top; r-- {
			res = append(res, matrix[r][left])
		}
		left++
	}

	if debugSpiralMatrix {
		fmt.Printf("spiralOrder output=%v\n", res)
	}
	return res
}
