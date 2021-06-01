package main

func singleNumber(nums []int) int {
	singleNum := 0
	for _, n := range nums {
		singleNum ^= n
	}

	return singleNum
}
