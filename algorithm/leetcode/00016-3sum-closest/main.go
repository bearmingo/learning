package main

import (
	"sort"
	"math"
	"fmt"
)

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	closest := math.MaxInt32
	ret := 0

	minS := nums[0] + nums[1] + nums[2]
	if minS > target {
		return minS
	}
	maxS := nums[len(nums) - 1] + nums[len(nums) - 2] + nums[len(nums) - 3]
	if maxS < target {
		return maxS
	}

	for k := 0; k < len(nums); k++ {
		first := nums[k]
		if k > 0 && first == nums[k - 1] {
			continue
		}
		
		for i, j := k + 1, len(nums) - 1; i < j; {
			ni := nums[i]
			nj := nums[j]
			sum := first + ni + nj
			diff := abs(target - sum)
			if diff == 0 {
				return sum
			} 

			if (diff < closest) {
				closest = diff
				ret = sum	
			}
			
			if (sum < target) {
				for i < j && nums[i] == nums[i + 1] { i++ }
				i++
			} else {
				for i < j && nums[j - 1] == nums[j] { j-- }
				j--
			}
		}
	}

	return ret
}

var ret = map[int]string{
    0: "",
    
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{1,1,1,1}, 0))

	fmt.Println(math.Exp)
}