package main

import "fmt"

// Leetcode
//    Runtime: 0ms, faster than 100.0%
//    Memory: 2.6MS, less than 57.69%

func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}

	// find out the index of the smallest element
	left, right := 0, len(nums)-1
	mid := 0
	for left < right {
		mid = (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	start := left

	// since we now know the start, find out if the target is
	// to left or right of start in the array.
	left, right = 0, len(nums)-1
	if target >= nums[start] && target <= nums[right] {
		left = start
	} else {
		right = start
	}

	for left <= right {
		mid = (left + right) / 2
		midNum := nums[mid]
		if midNum == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{5, 1, 3}, 1))
}
