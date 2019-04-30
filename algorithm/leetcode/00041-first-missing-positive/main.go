package main

// Leetcode
// https://leetcode.com/problems/first-missing-positive/
//
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.2 MB, less than 100.00%

import "fmt"

func firstMissingPositive(nums []int) int {
	// 先将数字放在对应的位置上
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] < len(nums) && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	// 查找第一个与编号不对应的值
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
