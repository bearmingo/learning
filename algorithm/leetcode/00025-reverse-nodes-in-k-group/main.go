package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	stack := make([]*ListNode, k)
	stackLen := 0

	var newHead, newEnd *ListNode

	for tmp := head; tmp != nil; tmp = tmp.Next {
		if stackLen == k {
			for i := stackLen - 1; i >= 0; i-- {
				if newHead == nil {
					newHead = stack[i]
					newEnd = newHead
					newEnd.Next = nil
				} else {
					newEnd.Next = stack[i]
					newEnd = newEnd.Next
					newEnd.Next = nil
				}
			}
			stackLen = 0
		}

		stack[stackLen] = tmp
		stackLen++
	}

	if stackLen == k {
		for i := stackLen - 1; i >= 0; i-- {
			if newHead == nil {
				newHead = stack[i]
				newEnd = newHead
				newEnd.Next = nil
			} else {
				newEnd.Next = stack[i]
				newEnd = newEnd.Next
				newEnd.Next = nil
			}
		}
	} else {
		for i := 0; i < stackLen; i++ {
			if newHead == nil {
				newHead = stack[i]
				newEnd = newHead
				newEnd.Next = nil
			} else {
				newEnd.Next = stack[i]
				newEnd = newEnd.Next
				newEnd.Next = nil
			}
		}
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
	// nodes := createListNode([]int{1, 2, 3, 4, 5})
	nodes := createListNode([]int{1})

	head := reverseKGroup(nodes, 1)

	printListVals(head)
}
