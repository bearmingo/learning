package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

type node struct {
	x, y int
	step int
}

func findNearestLand(grid [][]int, x, y int) int {
	m, n := len(grid), len(grid[0])
	flag := make([]bool, n*m)

	queue := make([]node, 0)
	queue = append(queue, node{x: x, y: y, step: 0})
	flag[x*m+y] = true

	for len(queue) != 0 {
		f := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			nx, ny := f.x+dx[i], f.y+dy[i]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}

			if !flag[nx*m+ny] {
				if grid[nx][ny] == 1 {
					return f.step + 1
				}

				queue = append(queue, node{x: nx, y: ny, step: f.step + 1})
				flag[nx*m+ny] = true
			}
		}
	}

	return -1
}

func maxDistance(grid [][]int) int {
	dis := -1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 0 {
				dis = max(findNearestLand(grid, i, j), dis)
			}
		}
	}
	return dis
}
