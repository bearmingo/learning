package main

// https://leetcode-cn.com/problems/distribute-coins-in-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dfs(node *TreeNode, cnt *int) int {
	if node == nil {
		return 0
	}

	l := dfs(node.Left, cnt)
	r := dfs(node.Right, cnt)
	*cnt += abs(l) + abs(r)

	return node.Val + l + r - 1
}

func distributeCoins(root *TreeNode) int {
	cnt := new(int)
	dfs(root, cnt)
	return *cnt
}
