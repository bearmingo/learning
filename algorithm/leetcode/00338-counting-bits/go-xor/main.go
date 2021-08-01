package main

func hammingWeight(num uint32) int {
	ret := 0
	for ; num > 0; num &= num - 1 {
		ret++
	}

	return ret
}

func countBits(n int) []int {
	rets := make([]int, 0, n)

	for i := 0; i <= n; i++ {
		rets = append(rets, hammingWeight(uint32(i)))
	}
	return rets
}
