package main

// Leetcode
// https://leetcode.com/problems/combination-sum-ii/
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 2.7 MB, less than 100.00%

import (
	"fmt"
	"sort"
)

// combinationNext Return true if stop search backward
func combinationNext(candidates, cache []int, results *[][]int, index, cacheSize, remainder int) {
	for i := index; i < len(candidates); i++ {
		t := candidates[i]
		r := remainder - t
		if r == 0 {
			found := make([]int, cacheSize, cacheSize+1)
			copy(found, cache[:cacheSize])
			found = append(found, t)
			*results = append(*results, found)
		} else if r > 0 && i+1 < len(candidates) {
			nextI := i + 1
			if nextI < len(candidates) {
				cache[cacheSize] = t
				combinationNext(candidates, cache, results, nextI, cacheSize+1, r)
			} else {
				return
			}
		} else if r < 0 {
			return
		}

		// 如果后面一个数字与当前的相同，搜索出来的结果会与当前的结果重复。
		// 所以直接跳过了。
		// 例如 [1,1,1,1,3] , target=4, 第一个1搜索过后后续1和4的说有组合
		// 都不需要再搜索了。
		for ; i+1 < len(candidates) && t == candidates[i+1]; i++ {
		}
	}
	return
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	results := make([][]int, 0)
	cache := make([]int, len(candidates))

	combinationNext(candidates, cache, &results, 0, 0, target)

	return results
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
