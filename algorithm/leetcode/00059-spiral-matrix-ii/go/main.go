package main

import "fmt"

func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	beginRow, endRow := 0, n-1
	beginCol, endCol := 0, n-1

	v := 1

	for beginRow <= endRow || beginCol <= endCol {
		// upper row: left -> right
		for i := beginCol; i <= endCol; i++ {
			ret[beginRow][i] = v
			v++
		}
		beginRow++

		// right col: up -> down
		for i := beginRow; i <= endRow; i++ {
			ret[i][endCol] = v
			v++
		}
		endCol--

		if beginRow < endRow {
			for i := endCol; i >= beginCol; i-- {
				ret[endRow][i] = v
				v++
			}
			endRow--
		}

		if beginCol < endCol {
			for i := endRow; i >= beginRow; i-- {
				ret[i][beginCol] = v
				v++
			}
			beginCol++
		}
	}

	return ret
}

func printResult(r [][]int) {
	for _, row := range r {
		fmt.Println(row)
	}
}

func main() {
	printResult(generateMatrix(0))
	printResult(generateMatrix(1))
	printResult(generateMatrix(4))
}
