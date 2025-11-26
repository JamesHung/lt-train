package partitionequalsubsetsum

const debugPartitionEqualSubsetSum = false

// // canPartition returns true if nums can be split into two subsets with equal sum.
// func canPartition(nums []int) bool {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}

// 	if total%2 != 0 { // odd sum cannot be split evenly
// 		return false
// 	}

// 	target := total / 2

// 	// 正向一維 DP：這一輪寫入 next，讀取的是上一輪 dp，避免重複使用同一元素。
// 	dp := make([]bool, target+1)
// 	dp[0] = true

// 	for _, num := range nums {
// 		next := make([]bool, target+1)
// 		copy(next, dp)
// 		for s := 0; s+num <= target; s++ {
// 			if dp[s] {
// 				next[s+num] = true
// 			}
// 		}
// 		dp = next
// 		if debugPartitionEqualSubsetSum {
// 			fmt.Printf("after num=%d dp=%v\n", num, dp)
// 		}
// 	}

// 	return dp[target]
// }

func canPartition(nums []int) bool {
	total := 0

	for _, num := range nums {
		total += num
	}

	if total%2 == 1 {
		return false
	}

	target := total / 2

	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		temp := make([]bool, target+1)
		copy(temp, dp)

		for i := 0; i+num <= target; i++ {
			if dp[i] {
				temp[i+num] = true
			}
		}
		dp = temp
	}

	return dp[target]
}
