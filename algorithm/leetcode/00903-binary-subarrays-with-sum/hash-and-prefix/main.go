package main

// 前缀和加哈希
func numSubarraysWithSum(A []int, S int) int {
	result := 0
	prefix := make(map[int]int, len(A))
	sum := 0
	for e := range A {
		sum += A[e]
		if sum == S {
			result++
		}
		result += prefix[sum-S]
		prefix[sum] = prefix[sum] + 1
	}
	return result
}
