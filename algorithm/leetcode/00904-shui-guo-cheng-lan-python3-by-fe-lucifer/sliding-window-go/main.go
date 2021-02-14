package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func totalFruit(tree []int) int {
	counter := make(map[int]int)
	for _, v := range tree {
		counter[v] = 0
	}

	basket := 2

	result := 0

	var left, right int
	for right = 0; right < len(tree); right++ {
		nR := tree[right]
		if counter[nR] == 0 {
			basket--
		}
		counter[nR]++

		for basket < 0 {
			nL := tree[left]
			counter[nL]--
			if counter[nL] == 0 {
				basket++
			}
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
