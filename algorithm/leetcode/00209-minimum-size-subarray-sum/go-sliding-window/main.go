package main

// INT_MAX Max value of int
const maxInt = int(^uint(0) >> 1)

func minSubArrayLen(s int, nums []int) int {
	var numLen = len(nums)
	if numLen == 0 {
		return 0
	}
	if nums[0] >= s {
		return 1
	}

	var minLen = maxInt
	var sum = 0
	var l = 0
	var r = 0

	for ; r < numLen; r++ {
		rnum := nums[r]
		if rnum >= s {
			return 1
		}
		sum += rnum

		for sum >= s {
			if r-l+1 < minLen {
				minLen = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}

	if minLen == maxInt {
		return 0
	}
	return minLen

}
