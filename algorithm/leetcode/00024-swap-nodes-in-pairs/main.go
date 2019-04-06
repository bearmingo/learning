package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// 1    3     4     5
// ^a   ^pre  ^p2
// 3    1     5     4     6     9
//      ^pre        ^p2

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	newHead := head.Next
	next := head

	next.Next = newHead.Next
	newHead.Next = next

	pre := next
	p2 := next.Next

	for p2 != nil && p2.Next != nil {
		pre.Next = p2.Next
		p2.Next = p2.Next.Next
		pre.Next.Next = p2

		pre = p2
		p2 = p2.Next
	}

	return newHead
}

// test code

func createListNode(nums []int) *ListNode {
	var head, end *ListNode

	for _, n := range nums {
		node := &ListNode{Val: n}
		if head == nil {
			head = node
			end = node
		} else {
			end.Next = node
			end = node
		}
	}
	return head
}

func printListVals(head *ListNode) {
	isFirst := true
	buf := bytes.Buffer{}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		if isFirst {
			isFirst = false
		} else {
			buf.WriteString("->")
		}
		buf.WriteString(strconv.Itoa(tmp.Val))
	}
	fmt.Println(buf.String())
}

func main() {
	nodes := createListNode([]int{1, 2, 3, 4, 5})

	head := swapPairs(nodes)

	printListVals(head)
}
