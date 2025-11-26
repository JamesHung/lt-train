package stringtointegeratoi

const (
	MAX_INT32 = 1<<31 - 1
	MIN_INT32 = -1 << 31
)

// myAtoi converts a string to a 32-bit signed integer following the described rules.
func myAtoi(s string) int {

	i, n := 0, len(s)

	for i < n && s[i] == ' ' {
		i++
	}

	if i == n {
		return 0
	}

	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	value := 0

	for i < n {
		char := s[i]

		if char < '0' || char > '9' {
			break
		}

		digit := int(char - '0')

		if sign == 1 {
			//這樣會orverflow
			// if value*10 + digit > MAX_INT32 {

			// }
			if value > (MAX_INT32-digit)/10 {
				return MAX_INT32
			}
		} else {
			// max = abs(min) -1
			// abs(min) = max+1
			// if -(value*10+dight) < -(max+1)
			if value > (MAX_INT32+1-digit)/10 {
				return MIN_INT32
			}
		}

		value = value*10 + digit
		i++
	}

	if sign == 1 {
		return value
	}

	return -value
}
