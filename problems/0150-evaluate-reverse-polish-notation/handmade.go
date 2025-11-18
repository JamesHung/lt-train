package evaluatorpn

import (
	"strconv"
)

const debugEvalRPN = false

// evalRPN evaluates the tokens using a stack.
func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, c := range tokens {
		switch c {
		case "+", "-", "*", "/":
			first := stack[len(stack)-1]
			second := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			newVal := operate(c, first, second)
			stack = append(stack, newVal)
		default:
			num, err := strconv.Atoi(c)
			if err != nil {
				continue
			}
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func operate(op string, first, second int) int {
	switch op {
	case "+":
		return second + first
	case "-":
		return second - first
	case "*":
		return second * first
	case "/":
		return second / first
	default:
		return -1
	}
}
