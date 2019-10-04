package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	left, top := 0, 0
	bottom := len(matrix) - 1
	right := len(matrix[0]) - 1

	result := make([]int, 0, len(matrix)*len(matrix[0]))

	for left <= right && top <= bottom {
		// top, left -> right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// right, top -> bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// bottom, right -> left, if row remains
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// left, bottom -> top, if column remains
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	test := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	test2 := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}

	fmt.Println(spiralOrder(test))
	fmt.Println(spiralOrder(test2))
}
