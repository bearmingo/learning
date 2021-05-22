package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestOnes(nums []int, k int) int {
	longestLen := 0

	left, right := 0, 0

	used := 0

	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			if used < k {
				nums[right] = 2
				used += 1
			} else {
				for left < right {
					if nums[left] != 2 {
						left++
						continue
					}
					nums[left] = 0
					nums[right] = 2
					left++
					break
				}
			}
		} else {
			for ; left < right; left++ {
				if nums[left] != 0 {
					break
				}
			}
		}

		if left != right {
			longestLen = max(right-left+1, longestLen)
		}
	}

	return longestLen
}
