package main

func countBits(n int) []int {
	bits := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}
		bits[i] = bits[i-highBit] + 1
	}
	return bits
}
