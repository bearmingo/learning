package main

type node struct {
	next  *node
	value int
}

type MyQueue struct {
	head *node
	tail *node
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	n := &node{value: x}
	if this.head == nil {
		this.head = n
	}
	if this.tail != nil {
		this.tail.next = n
	}
	this.tail = n
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.head == nil {
		return 0
	}

	n := this.head

	if n == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.head = this.head.next
	}

	return n.value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.head == nil {
		return 0
	}
	return this.head.value
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.head == nil
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
