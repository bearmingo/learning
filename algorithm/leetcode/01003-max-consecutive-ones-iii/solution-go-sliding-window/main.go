package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestOnes(nums []int, k int) int {
	var rsum, lsum, result int
	var left int
	for right, v := range nums {
		rsum += 1 - v
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
