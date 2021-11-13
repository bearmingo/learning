package main

func numSubarraysWithSum(A []int, S int) int {
	var result int
	var count int
	i, j := 0, 0

	for ; count < S && j < len(A); j++ {
		if A[j] == 1 {
			count++
		}
	}
	for j <= len(A) && count == S {
		left := 1
		for i < len(A) && A[i] == 0 {
			i++
			left++
		}
		right := 1
		for j < len(A) && A[j] == 0 {
			j++
			right++
		}
		if i == j {
			// S == 0的情况
			result += (left-1)*(left-2)/2 + left - 1
		} else {
			result += left * right
		}
		i++
		j++
	}

	return result
}
