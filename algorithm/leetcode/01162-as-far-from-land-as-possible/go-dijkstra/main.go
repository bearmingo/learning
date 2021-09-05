package main

import "math"

func maxDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	q := NewProirityQueye()

	d := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		d[i] = math.MaxInt32
	}

	for i := 0; i < m; i++ {
		tmp := i * m
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				d[tmp+j] = 0
				q.Push(status{v: 0, x: i, y: j})
			}
		}
	}

	for !q.Empty() {
		f := q.Top()
		q.Pop()

		for i := 0; i < 4; i++ {
			nx, ny := f.x+dx[i], f.y+dy[i]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}

			indx := nx*m + ny
			if f.v+1 < d[indx] {
				d[indx] = f.v + 1
				q.Push(status{v: d[indx], x: nx, y: ny})
			}
		}
	}

	ans := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans = max(ans, d[i*m+j])
			}
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

type status struct {
	v, x, y int
}

type priorityQueue struct {
	q []status
}

func NewProirityQueye() priorityQueue {
	return priorityQueue{q: make([]status, 0)}
}

// 从小到大排序
func (p *priorityQueue) Push(s status) {
	for i, n := range p.q {
		if n.v > s.v {
			p.q = append(append(p.q[0:i], s), p.q[i:]...)
			return
		}
	}

	p.q = append(p.q, s)
}

func (p *priorityQueue) Empty() bool {
	return len(p.q) == 0
}

func (p *priorityQueue) Top() status {
	return p.q[0]
}

func (p *priorityQueue) Pop() {
	p.q = p.q[1:]
}

var (
	dx = []int{-1, 0, 1, 0}
	dy = []int{0, 1, 0, -1}
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
