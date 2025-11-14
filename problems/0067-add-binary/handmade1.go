package addbinary

// addBinaryReverse demonstrates the reversing approach.
func addBinaryReverse(a, b string) string {
	carry := 0
	aLength := len(a) - 1
	bLength := len(b) - 1
	i, j := aLength, bLength
	// result := strings.Builder{}
	result := []byte{}

	for i >= 0 || j >= 0 {
		var aDigit, bDigit int
		if i >= 0 {
			aDigit = int(a[i] - '0')
			i--
		}

		if j >= 0 {
			bDigit = int(b[j] - '0')
			j--
		}

		sum := aDigit + bDigit + carry
		carry = sum / 2
		result = append(result, byte((sum%2)+'0'))
		// result.WriteByte(byte((sum % 2) + '0'))
	}

	if carry > 0 {
		// result.WriteByte(byte(carry + '0'))
		result = append(result, byte(carry+'0'))
	}

	reverseBytes(result)
	return string(result)
}

func reverseBytes(bytes []byte) {
	lo, hi := 0, len(bytes)-1

	for lo < hi {
		bytes[lo], bytes[hi] = bytes[hi], bytes[lo]
		lo++
		hi--
	}
}
