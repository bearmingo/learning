package main

import "container/list"

type FreqStack struct {
	freq    map[int]int
	group   map[int]*list.List
	maxFreq int
}

func Constructor() FreqStack {
	return FreqStack{
		freq:  make(map[int]int),
		group: make(map[int]*list.List),
	}
}

func (this *FreqStack) Push(val int) {
	f, ok := this.freq[val]
	if !ok {
		f = 0
	}
	f += 1
	this.freq[val] = f
	if this.maxFreq < f {
		this.maxFreq = f
	}

	g, ok := this.group[f]
	if !ok {
		g = list.New()
		this.group[f] = g
	}
	g.PushBack(val)
}

func (this *FreqStack) Pop() int {
	g := this.group[this.maxFreq]
	e := g.Back()
	g.Remove(e)

	v := e.Value.(int)
	this.freq[v] -= 1
	if g.Len() == 0 {
		delete(this.group, this.maxFreq)
		this.maxFreq -= 1
	}
	return v
}
