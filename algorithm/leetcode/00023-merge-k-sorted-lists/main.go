package main

// 使用最小堆得排序算法。
//

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type nodeHeap struct {
	data []*ListNode
	len  int
}

// Override heap interface

func (m *nodeHeap) Len() int {
	return m.len
}

func (m *nodeHeap) Less(i, j int) bool {
	return m.data[i].Val < m.data[j].Val
}

func (m *nodeHeap) Swap(i, j int) {
	if i < 0 || i > m.len || j < 0 || j > m.len {
		return
	}
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *nodeHeap) Push(x interface{}) {
	if x == nil {
		return
	}
	m.data[m.len] = x.(*ListNode)
	m.len++
}

func (m *nodeHeap) Pop() interface{} {
	if m.len == 0 {
		return nil
	}

	item := m.data[m.len-1]
	m.len--
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	nh := &nodeHeap{
		data: make([]*ListNode, len(lists)),
		len:  0,
	}
	heap.Init(nh)

	for _, item := range lists {
		if item == nil {
			continue
		}
		heap.Push(nh, item)
	}

	if nh.Len() == 0 {
		return nil
	}

	head := heap.Pop(nh).(*ListNode)
	end := head
	if end.Next != nil {
		heap.Push(nh, end.Next)
	}

	next, ok := (heap.Pop(nh)).(*ListNode)
	for ok {
		if end.Next != next {
			end.Next = next
		}
		end = next
		if next.Next != nil {
			heap.Push(nh, next.Next)
		}

		if nh.len == 0 {
			break
		}

		next, ok = (heap.Pop(nh)).(*ListNode)
	}

	return head
}

func main() {
	lists := make([]*ListNode, 3)

	/*
		lists[0] = &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		}
		lists[1] = &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		}
		lists[2] = &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		}
	*/

	mergeKLists(lists)
}
