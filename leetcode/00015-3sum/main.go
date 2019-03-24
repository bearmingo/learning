package main

import (
	"sort"
	"fmt"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	ret := make([][]int, 0)

	for k := 0; k < len(nums); k++ {
		first := nums[k]
		if first > 0 {
			break
		}
		if k > 0 && first == nums[k - 1] {
			continue
		}
		
		for i, j := k + 1, len(nums) - 1; i < j; {
			ni := nums[i]
			nj := nums[j]
			sum := first + ni + nj
			if sum == 0 {
				ret = append(ret, []int{first, ni, nj})

				for i < j && nums[i] == nums[i + 1] { i++ }
				for i < j && nums[j - 1] == nums[j] { j-- }
				i++
				j--
			} else if (sum < 0) {
				i++
			} else {
				j--
			}
		}

	}

	return ret
}

func main() {
	fmt.Print(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}