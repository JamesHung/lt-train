package productoftwoexceptself

import "fmt"

const debugProductExceptSelf = false

// productExceptSelf returns array where each entry is product of all other elements.
func productExceptSelf(nums []int) []int {
	// n := len(nums)
	// out := make([]int, n)

	// prefix := 1
	// for i := 0; i < n; i++ {
	// 	out[i] = prefix
	// 	prefix *= nums[i]
	// 	if debugProductExceptSelf {
	// 		fmt.Printf("prefix[%d]=%d\n", i, out[i])
	// 	}
	// }

	// suffix := 1
	// for i := n - 1; i >= 0; i-- {
	// 	out[i] *= suffix
	// 	if debugProductExceptSelf {
	// 		fmt.Printf("suffix step idx=%d -> %d\n", i, out[i])
	// 	}
	// 	suffix *= nums[i]
	// }

	// return out
	n := len(nums)
	out := make([]int, n)
	prefix := make([]int, n)
	postfix := make([]int, n)

	prefix[0] = 1
	if debugProductExceptSelf {
		fmt.Printf("prefix[%d]=%d\n", 0, prefix[0])
	}
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
		if debugProductExceptSelf {
			fmt.Printf("prefix[%d]=%d\n", i, prefix[i])
		}
	}

	postfix[n-1] = 1
	if debugProductExceptSelf {
		fmt.Printf("postfix[%d]=%d\n", n-1, postfix[n-1])
	}
	for i := n - 2; i >= 0; i-- {
		postfix[i] = postfix[i+1] * nums[i+1]
		if debugProductExceptSelf {
			fmt.Printf("postfix[%d]=%d\n", i, postfix[i])
		}
	}

	for i := 0; i < n; i++ {
		out[i] = prefix[i] * postfix[i]
		if debugProductExceptSelf {
			fmt.Printf("out[%d]=%d\n", i, out[i])
		}
	}

	return out
}
