package main

import "math"

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	f := make([][]int, m)
	for i := 0; i < m; i++ {
		c := make([]int, n)
		f[i] = c

		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				c[j] = 0
			} else {
				c[j] = math.MaxInt32
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}

			if i-1 >= 0 {
				f[i][j] = min(f[i][j], f[i-1][j]+1)
			}
			if j-1 >= 0 {
				f[i][j] = min(f[i][j], f[i][j-1]+1)
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}

			if i+1 < n {
				f[i][j] = min(f[i][j], f[i+1][j]+1)
			}
			if j+1 < n {
				f[i][j] = min(f[i][j], f[i][j+1]+1)
			}
		}
	}

	ans := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans = max(ans, f[i][j])
			}
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
