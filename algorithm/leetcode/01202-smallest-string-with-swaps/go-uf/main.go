package main

import (
	"sort"
	"strings"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	u := New(len(s))
	for _, i := range pairs {
		u.Union(i[0], i[1])
	}

	chaMap := make(map[int][]rune)
	for i, c := range s {
		p := u.Find(i)
		if v, ok := chaMap[p]; ok {
			chaMap[p] = append(v, c)
		} else {
			v := make([]rune, 0)
			v = append(v, c)
			chaMap[p] = v
		}
	}

	for _, v := range chaMap {
		sort.Sort(runeSlice(v))
	}

	b := strings.Builder{}
	for i := 0; i < len(s); i++ {
		r := u.Find(i)
		chs := chaMap[r]
		b.WriteRune(chs[0])
		chaMap[r] = chs[1:]
	}

	return b.String()
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

type runeSlice []rune

func (r runeSlice) Len() int {
	return len(r)
}

func (r runeSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func (r runeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
