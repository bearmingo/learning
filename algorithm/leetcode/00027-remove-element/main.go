package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	len := 0
	for i, num := range nums {
		if num == val {
			continue
		}

		if len != i {
			nums[len] = num
		}
		len++
	}

	return len
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	len := removeElement(nums, 3)

	for i := 0; i < len; i++ {
		fmt.Print(nums[i], " ")
	}
}
