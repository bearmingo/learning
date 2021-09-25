package main

func findCircleNum(isConnected [][]int) int {
	if isConnected == nil {
		return 0
	}

	n := len(isConnected)

	u := New(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				u.Union(i, j)
			}
		}
	}

	return u.Size()
}

type unionFind struct {
	parent []int
	size   int
}

func New(n int) *unionFind {
	u := &unionFind{
		parent: make([]int, n),
		size:   n,
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (u *unionFind) Find(x int) int {
	r := x
	for ; r != u.parent[r]; r = u.parent[r] {
	}

	for k := x; k != r; {
		j := u.parent[k]
		u.parent[k] = r
		k = j
	}
	return r
}

func (u *unionFind) Union(p, q int) {
	pp, pq := u.Find(p), u.Find(q)
	if pp == pq {
		return
	}
	u.parent[pp] = pq
	u.size--
}

func (u *unionFind) IsConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *unionFind) Size() int {
	return u.size
}
