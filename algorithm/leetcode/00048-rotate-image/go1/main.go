package main

// 0, 0 -> 0, 2
// 0, 1 -> 1, 2
// 0, 2 -> 2, 2
// 1, 0 -> 0, 1
// 1, 1 -> 1, 1
// 1, 2 -> 2, 1
// 2, 0 -> 0, 0
// 2, 1 -> 1, 0
// 2, 2 -> 2, 0

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	h := len(matrix)
	w := len(matrix[0])

	ret := make([][]int, h)
	for i := 0; i < h; i++ {
		tmp := make([]int, w)
		copy(tmp, matrix[i])
		ret[i] = tmp
	}

	for i, vals := range ret {
		for j, v := range vals {
			matrix[j][w-1-i] = v
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	print(matrix)
}
