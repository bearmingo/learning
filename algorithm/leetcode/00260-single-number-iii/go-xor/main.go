package main

func singleNumber(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	rightmostbit := xor & -xor

	a, b := 0, 0

	for _, n := range nums {
		if n&rightmostbit == 0 {
			a ^= n
		} else {
			b ^= n
		}
	}

	return []int{a, b}
}
