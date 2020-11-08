package main

import "fmt"

// https://leetcode-cn.com/problems/jump-game-ii/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	var (
		n        int = len(nums)
		cnt      int = 0
		furthest int = 0
		end      int = 0
	)

	for i := 0; i < n-1; i++ {
		s := nums[i]
		furthest = max(furthest, s+i)
		if i == end {
			cnt = cnt + 1
			end = furthest
		}
	}

	return cnt
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
