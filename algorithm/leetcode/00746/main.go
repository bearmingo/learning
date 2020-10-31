package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	var ppCost, pCost, now int

	for _, c := range cost {
		now = min(ppCost, pCost) + c

		ppCost = pCost
		pCost = now
	}

	return min(pCost, ppCost)
}

func main() {
	fmt.Print(minCostClimbingStairs([]int{10, 15, 20}), "\n")
	fmt.Print(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
