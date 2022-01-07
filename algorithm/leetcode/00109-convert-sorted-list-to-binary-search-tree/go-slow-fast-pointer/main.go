package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return makeTree(head)
}

func makeTree(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	slow, fast, prev := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	for prev.Next != slow {
		prev = prev.Next
	}

	root := &TreeNode{Val: slow.Val}
	prev.Next = nil
	rightHead := slow.Next
	root.Left = makeTree(head)
	root.Right = makeTree(rightHead)

	prev.Next = slow
	return root
}
