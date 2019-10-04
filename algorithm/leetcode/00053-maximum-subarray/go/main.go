package main

import "fmt"

func maxSubArray(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < l; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
