package spiralmatrix

const debugSpiralMatrix = false

// spiralOrder returns all elements of the matrix in spiral order.
func spiralOrder(matrix [][]int) []int {
	rowLength := len(matrix)
	colLength := len(matrix[0])

	result := make([]int, 0, rowLength*colLength)

	left, right, top, bottom := 0, colLength-1, 0, rowLength-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		for j := top; j <= bottom; j++ {
			result = append(result, matrix[j][right])
		}

		right--
		if left > right {
			break
		}

		for k := right; k >= left; k-- {
			result = append(result, matrix[bottom][k])
		}
		bottom--
		if left > right {
			break
		}

		for l := bottom; l >= top; l-- {
			result = append(result, matrix[l][left])
		}
		left++
		if top > bottom {
			break
		}

	}

	return result
}
