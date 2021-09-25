package main

func equationsPossible(equations []string) bool {
	unionFind := New(26)

	for _, item := range equations {
		if item[1] == '=' {
			l := int(item[0]) - 'a'
			r := int(item[3]) - 'a'
			unionFind.Union(l, r)
		}
	}
	for _, item := range equations {
		if item[1] == '!' {
			l := int(item[0]) - 'a'
			r := int(item[3]) - 'a'
			if unionFind.IsConnected(l, r) {
				return false
			}
		}
	}

	return true
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
