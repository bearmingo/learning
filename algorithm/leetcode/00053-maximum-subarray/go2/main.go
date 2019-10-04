package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	curMax := nums[0]
	acc := nums[0]

	for _, n := range nums[1:] {
		acc = max(acc+n, n)
		if acc > curMax {
			curMax = acc
		}
	}

	return curMax
}

func main() {
	fmt.Println(maxSubArray([]int{1}))
}
