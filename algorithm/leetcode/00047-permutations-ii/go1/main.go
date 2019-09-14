package main

import (
	"fmt"
	"strconv"
)

func makeKey(nums []int) string {
	var key = ""
	for _, n := range nums {
		key += strconv.Itoa(n) + ","
	}
	return key
}

func permuteRecursive(nums []int, begin int, ret map[string][]int) {
	if begin >= len(nums) {
		key := makeKey(nums)
		if _, ok := ret[key]; !ok {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ret[key] = temp
		}
	}

	for i := begin; i < len(nums); i++ {
		nums[i], nums[begin] = nums[begin], nums[i]
		permuteRecursive(nums, begin+1, ret)
		nums[begin], nums[i] = nums[i], nums[begin]
	}
}

func permuteUnique(nums []int) [][]int {
	ret := make(map[string][]int)
	permuteRecursive(nums, 0, ret)

	r := make([][]int, 0, len(ret))
	for _, value := range ret {
		r = append(r, value)
	}

	return r
}

func main() {
	fmt.Print(permuteUnique([]int{1, 1, 2}))
}
