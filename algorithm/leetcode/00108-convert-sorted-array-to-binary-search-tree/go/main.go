package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return makeTree(nums, 0, len(nums)-1)
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
