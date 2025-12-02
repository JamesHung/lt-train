package containerwithmostwater

const debugContainer = false

// maxArea finds the largest area formed by two lines.
// Two-pointer approach: move the shorter side inward each step.
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	best := 0
	for left <= right {
		volumn := min(height[left], height[right]) * (right - left)

		if volumn > best {
			best = volumn
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return best
}
