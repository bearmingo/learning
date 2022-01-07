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
	l := make([]int, 0)
	for i := head; i != nil; i = i.Next {
		l = append(l, i.Val)
	}
	return makeTree(l, 0, len(l)-1)
}

func makeTree(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = makeTree(nums, left, mid-1)
	root.Right = makeTree(nums, mid+1, right)
	return root
}
