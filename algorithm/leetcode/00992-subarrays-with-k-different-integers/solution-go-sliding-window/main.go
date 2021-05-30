package main

func atMostDistance(nums []int, k int) int {
	left := 0

	// 记录每一个数字出现的次数
	freq := make(map[int]int)
	// 记录有几个数字出现过
	count := 0
	result := 0

	for right, n := range nums {
		if freq[n] == 0 {
			count += 1
		}
		freq[n]++

		for count > k {
			lnum := nums[left]
			freq[lnum]--
			if freq[lnum] == 0 {
				count--
			}

			left++
		}
		result += right - left + 1
	}

	return result
}

func subarraysWithKDistinct(nums []int, k int) int {
	return atMostDistance(nums, k) - atMostDistance(nums, k-1)
}
