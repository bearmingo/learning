package main

import (
	"fmt"
	"sort"
)

func dfs(nums []int, used []bool, l []int, ret *[][]int) {
	if len(l) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, l)
		*ret = append(*ret, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		used[i] = true
		temp := append(l, nums[i])
		dfs(nums, used, temp, ret)
		used[i] = false
	}
}

func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	used := make([]bool, len(nums))

	l := make([]int, 0)
	sort.Ints(nums)

	dfs(nums, used, l, &ret)

	return ret
}

func main() {
	fmt.Print(permuteUnique([]int{1, 1, 2}))
}
