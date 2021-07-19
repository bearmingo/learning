package main

func findErrorNums(nums []int) []int {
	xor := 0
	for i := 1; i <= len(nums); i++ {
		xor ^= i
	}
	for _, n := range nums {
		xor ^= n
	}

	// xor = a^b

	// 第一位不是0,
	// rightmostbit := 1
	// for ; (xor & rightmostbit) == 0; rightmostbit <<= 1 {
	// }
	rightmostbit := xor & -xor

	a, b := 0, 0
	for _, n := range nums {
		// 根据该位是否为0分为两组
		if (rightmostbit & n) == 0 {
			a ^= n
		} else {
			b ^= n
		}
	}

	for i := 1; i <= len(nums); i++ {
		if (rightmostbit & i) == 0 {
			a ^= i
		} else {
			b ^= i
		}
	}

	for _, n := range nums {
		if n == a {
			return []int{a, b}
		}
	}

	return []int{b, a}
}
