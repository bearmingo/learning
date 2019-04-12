package main

import "fmt"

// 全排列,升序
//
// LeetCode
//   Runtime: 4ms, faster than 100%
//   Memory: 3.4MB, less than 16.7%

// 反转字符串
func reverse(nums []int, begin, end int) {
	for from, to := begin, end-1; from < to; from, to = from+1, to-1 {
		nums[from], nums[to] = nums[to], nums[from]
	}
}

func nextPermutation(nums []int) {
	// Special cases
	if len(nums) <= 1 {
		return
	}

	// find i and j, which num[i] < num[j]
	i, j := len(nums)-2, len(nums)-1
	for ; i >= 0; i, j = i-1, j-1 {
		if nums[i] < nums[j] {
			break
		}
	}

	if i < 0 {
		reverse(nums, 0, len(nums))
		return
	}

	// find k
	k := len(nums) - 1
	for ; k >= j; k-- {
		if nums[i] < nums[k] {
			break
		}
	}

	// swap i and k
	nums[i], nums[k] = nums[k], nums[i]
	reverse(nums, j, len(nums))
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
