package main

type LFUCache struct {
	cap   int
	datas map[int]*node

	freqLink *freqLinkNode
}

func (this *LFUCache) Get(key int) int {
	if d, ok := this.datas[key]; ok {
		incrFreq(d)
		return d.value
	}

	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	if d, ok := this.datas[key]; ok {
		d.value = value
		incrFreq(d)
		return
	}

	if this.cap == len(this.datas) {
		delData := this.freqLink.next.header.pre
		removeFromLink(delData)
		delData.freqLink = nil
		if this.freqLink.next.isEmpty() {
			removeFromFreqLink(this.freqLink.next)
		}

		delete(this.datas, delData.key)
	}

	newData := &node{key: key, value: value}
	this.datas[key] = newData

	freqLink := this.freqLink.next
	if freqLink.freq != 1 {
		freqLink = newFreqLinkNode(1)
		insertNextToFreqLink(freqLink, this.freqLink)
	}

	insertNodeNext(newData, freqLink.header)
	newData.freqLink = freqLink
}

func incrFreq(d *node) {
	lastFreqLink := d.freqLink
	removeFromLink(d)
	d.freqLink = nil

	var newFreqLink *freqLinkNode
	// 检查是否有下一个
	if lastFreqLink.next.freq == lastFreqLink.freq+1 {
		newFreqLink = lastFreqLink.next
	} else {
		newFreqLink = newFreqLinkNode(lastFreqLink.freq + 1)
		insertNextToFreqLink(newFreqLink, lastFreqLink)
	}
	insertNodeNext(d, newFreqLink.header)
	d.freqLink = newFreqLink

	if lastFreqLink.isEmpty() {
		removeFromFreqLink(lastFreqLink)
		lastFreqLink.header.pre = nil
		lastFreqLink.header.next = nil
	}
}

type node struct {
	key   int
	value int

	pre      *node
	next     *node
	freqLink *freqLinkNode
}

type freqLinkNode struct {
	freq   int
	pre    *freqLinkNode
	next   *freqLinkNode
	header *node
}

func newFreqLinkNode(freq int) *freqLinkNode {
	d := &node{value: -1}
	d.pre = d
	d.next = d
	return &freqLinkNode{freq: freq, header: d}
}

func (f freqLinkNode) isEmpty() bool {
	return f.header.pre == f.header
}

func Constructor(capacity int) LFUCache {
	freq := &freqLinkNode{freq: -1}
	freq.next = freq
	freq.pre = freq
	c := LFUCache{cap: capacity, datas: make(map[int]*node), freqLink: freq}
	return c
}

func insertNodeNext(n, nextTo *node) {
	nextTo.next.pre = n
	n.next = nextTo.next
	nextTo.next = n
	n.pre = nextTo
}

func removeFromLink(b *node) {
	b.pre.next = b.next
	b.next.pre = b.pre
	b.next = nil
	b.pre = nil
}

func insertNextToFreqLink(n, nextTo *freqLinkNode) {
	nextTo.next.pre = n
	n.next = nextTo.next
	nextTo.next = n
	n.pre = nextTo
}

func removeFromFreqLink(b *freqLinkNode) {
	b.pre.next = b.next
	b.next.pre = b.pre

	b.next = nil
	b.pre = nil
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
