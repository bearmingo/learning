package main

// 前缀和 + 哈希
// prefix: 和为x的的个数。因为数组中只保护0和1，和为x的个数的另一层含义是：
// sum: 数组中x_0 ~ x_i的和

func numSubarraysWithSum(nums []int, goal int) int {

	count := 0
	prefix := make(map[int]int, len(nums))
	sum := 0
	for _, v := range nums {
		sum += v
		if sum == goal {
			count++
		}

		count += prefix[sum-goal]
		prefix[sum] = prefix[sum] + 1
	}

	return count
}
