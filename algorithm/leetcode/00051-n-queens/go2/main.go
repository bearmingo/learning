package main

import "fmt"

func solve(result *[][]string, queens [][]byte, flagC, flag45, flag135 []bool, row, n int) {
	if row == n {
		one := make([]string, 0)
		for _, r := range queens {
			one = append(one, string(r))
		}
		*result = append(*result, one)
		return
	}

	for col := 0; col < n; col++ {
		if flagC[col] && flag45[row+col] && flag135[n-1+col-row] {
			flagC[col] = false
			flag45[row+col] = false
			flag135[n-1+col-row] = false
			queens[row][col] = 'Q'
			solve(result, queens, flagC, flag45, flag135, row+1, n)
			queens[row][col] = '.'
			flagC[col] = true
			flag45[row+col] = true
			flag135[n-1+col-row] = true
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

	flagC := make([]bool, n)
	for i := 0; i < n; i++ {
		flagC[i] = true
	}
	flag45 := make([]bool, 2*n-1)
	flag135 := make([]bool, 2*n-1)
	for i := 0; i < 2*n-1; i++ {
		flag45[i] = true
		flag135[i] = true
	}

	solve(&results, queens, flagC, flag45, flag135, 0, n)

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
