package main

func numSubarraysWithSum(nums []int, goal int) int {
	sum := 0
	count := 0
	i, j := 0, 0

	for ; sum < goal && j < len(nums); j++ {
		if nums[j] == 1 {
			sum++
		}
	}

	for j <= len(nums) && sum == goal {
		left := 1
		for i < len(nums) && nums[i] == 0 {
			i++
			left++
		}
		right := 1
		for j < len(nums) && nums[j] == 0 {
			j++
			right++
		}
		if i == j {
			count += (left-1)*(left-2)/2 + left - 1
		} else {
			count += left * right
		}

		i++
		j++
	}

	return count

}
