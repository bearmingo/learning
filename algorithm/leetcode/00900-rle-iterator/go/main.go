package main

type RLEIterator struct {
	encoding  []int
	lastIndex int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding: encoding}
}

func (this *RLEIterator) Next(n int) int {
	if this.lastIndex >= len(this.encoding) {
		return -1
	}

	for i := this.lastIndex; i < len(this.encoding); i += 2 {
		num := this.encoding[i]
		if num >= n {
			this.encoding[i] -= n
			this.lastIndex = i
			return this.encoding[i+1]
		} else {
			n -= num
		}
	}
	this.lastIndex = len(this.encoding)
	return -1
}
