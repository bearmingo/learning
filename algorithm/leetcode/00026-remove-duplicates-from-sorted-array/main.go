package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	pre := nums[0]
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == pre {
			continue
		} else {
			pre = nums[i]
			index++
			nums[index] = nums[i]

		}
	}
	index++
	return index
}

func main() {
	// nodes := createListNode([]int{1, 2, 3, 4, 5})
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len := removeDuplicates(nums)

	fmt.Println(len)
}
