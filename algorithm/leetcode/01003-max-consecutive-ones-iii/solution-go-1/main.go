package main

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		right++
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}

	}
	return right - left
}
