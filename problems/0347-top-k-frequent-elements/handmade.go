package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for count := len(nums); count >= 1; count-- {
		for _, num := range buckets[count] {
			result = append(result, num)
		}

		if len(result) >= k {
			break
		}
	}

	return result[:k]
}
