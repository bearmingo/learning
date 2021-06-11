package main

import "math"

func singleNumber(nums []int) int {
	var ret int

	for i := 0; i < 32; i++ {
		cnt := 0
		mask := 1 << i
		for _, n := range nums {
			if n&mask != 0 {
				cnt++
			}
		}
		if cnt%3 != 0 {
			ret |= mask
		}
	}

	if ret > math.MaxInt32 {
		return ret - math.MaxInt32*2 - 2
	} else {
		return ret
	}
}
