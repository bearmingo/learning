package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {

	emailToName := make(map[string]string)
	emailToID := make(map[string]int)
	u := NewUnionFind(len(accounts))

	for i, a := range accounts {
		accountName := a[0]

		for _, e := range a[1:] {
			if id, ok := emailToID[e]; ok {
				u.Union(i, id)
			}
			emailToID[e] = i
			emailToName[e] = accountName
		}
	}

	tempResult := make(map[int][]string)
	for k, v := range emailToID {
		p := u.Find(v)
		if r, ok := tempResult[p]; ok {
			r = append(r, k)
			tempResult[p] = r
		} else {
			r := make([]string, 0)
			r = append(r, k)
			tempResult[p] = r
		}
	}

	result := make([][]string, 0, len(tempResult))
	for _, v := range tempResult {
		sort.Strings(v)
		item := make([]string, 0, len(v)+1)
		item = append(append(item, emailToName[v[0]]), v...)
		result = append(result, item)
	}

	return result
}

type unionFind struct {
	parent []int
}

func NewUnionFind(n int) *unionFind {
	u := &unionFind{
		parent: make([]int, n),
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
}

func (u *unionFind) IsConnected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
