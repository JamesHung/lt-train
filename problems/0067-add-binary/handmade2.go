package addbinary

// addBinary is intentionally left blank for practice.
// This variant builds the result from the back and avoids reversing.
func addBinary(a, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	res := make([]byte, max(len(a), len(b))+1)
	idx := len(res) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		res[idx] = byte(sum%2) + '0'
		idx--
		carry = sum / 2
	}

	return string(res[idx+1:])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
