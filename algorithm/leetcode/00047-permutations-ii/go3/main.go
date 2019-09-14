package main

import (
	"fmt"
	"sort"
)

func backtrack(now, candidates []int, ret *[][]int) {
	if len(candidates) == 0 {
		tmp := make([]int, len(now))
		copy(tmp, now)
		*ret = append(*ret, tmp)
		return
	}

	for i, val := range candidates {
		if i != 0 && val == candidates[i-1] {
			continue
		}
		can := make([]int, len(candidates))
		copy(can, candidates)
		if i != len(candidates)-1 {
			backtrack(append(now, val), append(can[:i], can[i+1:]...), ret)
		} else {
			backtrack(append(now, val), can[:i], ret)
		}
	}
}

func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	sort.Ints(nums)
	backtrack([]int{}, nums, &ret)
	return ret
}

func main() {
	fmt.Print(permuteUnique([]int{1, 1, 2}))
}
