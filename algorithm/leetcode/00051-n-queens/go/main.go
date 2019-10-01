package main

import "fmt"

func isValid(queens [][]byte, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if queens[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queens[i][j] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if queens[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func solve(res *[][]string, queens [][]byte, row, n int) {
	if row == n {
		one := make([]string, 0)
		for _, r := range queens {
			one = append(one, string(r))
		}
		*res = append(*res, one)
		return
	}

	for col := 0; col != n; col++ {
		if isValid(queens, row, col, n) {
			queens[row][col] = byte('Q')
			solve(res, queens, row+1, n)
			queens[row][col] = '.'
		}
	}
}

func solveNQueens(n int) [][]string {
	results := make([][]string, 0)
	queens := make([][]byte, n)

	sample := make([]byte, n)
	for i := 0; i < n; i++ {
		sample[i] = '.'
	}

	for i := 0; i < n; i++ {
		row := make([]byte, n)
		copy(row, sample)
		queens[i] = row
	}

	solve(&results, queens, 0, n)

	return results
}

func main() {

	results := solveNQueens(4)
	for _, s := range results {
		for _, r := range s {
			fmt.Println(r)
		}
		fmt.Println("=================")
	}
}
