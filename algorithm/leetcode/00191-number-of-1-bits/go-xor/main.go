package main

func hammingWeight(num uint32) int {
	ret := 0
	for ; num > 0; num &= num - 1 {
		ret++
	}

	return ret
}
