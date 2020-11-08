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
		cnt      int = 0 // 已经跳跃的次数
		furthest int = 0 // 已遍历的数据中，最远可到达的位置
		end      int = 0 // 上次跳跃最后能到达的位置，与i相等时，必须跳跃一次
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
