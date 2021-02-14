package main

func totalFruit(tree []int) int {
	result := 0
	b1, b2 := -1, -1
	i := 0
	for ; i < len(tree); i++ {
		if b1 == -1 {
			b1 = i
			continue
		}
		if tree[b1] == tree[i] {
			continue
		}
		if b2 == -1 {
			b2 = i
			continue
		}
		if tree[b2] == tree[i] {
			continue
		}
		if i-b1 > result {
			result = i - b1
		}
		b1 = i - 1
		for ; b1 >= 0 && tree[b1] == tree[b1-1]; b1-- {
		}
		b2 = i
	}
	if i-b1 > result {
		return i - b1
	}
	return result
}
