package main

// https://leetcode.com/problems/4sum/

import (
	"fmt"
	"sort"
)

func findNSum(nums []int, N, target int, before []int, results *[][]int) {
	if len(nums) < N || N < 2 {
		return
	}

	if N == 2 {
		l, r := 0, len(nums)-1
		for l < r {
			s := nums[l] + nums[r]
			if s == target {
				found := append(before, nums[l], nums[r])
				*results = append(*results, found)

				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if s < target {
				l++
			} else {
				r--
			}
		}
	} else {
		// recursively reduce N
		for i := 0; i < len(nums)-N+1; i++ {
			// sorted list, nums[i] is minimus num,
			// any nums's sum will bigger than target < nums[i] * N
			if target < nums[i]*N || target > nums[len(nums)-1]*N {
				break
			}
			if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
				temp := append(before, nums[i])
				findNSum(nums[i+1:], N-1, target-nums[i], temp, results)
			}
		}
	}
}

func fourSum(nums []int, target int) [][]int {
	sortedNums := nums
	sort.Ints(sortedNums)

	results := make([][]int, 0, 10)

	findNSum(sortedNums, 4, target, make([]int, 0, 4), &results)

	return results
}

func main() {
	fmt.Print(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
