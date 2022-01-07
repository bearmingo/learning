package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(dfs(root.Left, 0)-dfs(root.Right, 0)) > 1 {
		return false
	}

	return isBalanced(root.Right) && isBalanced(root.Left)
}

func dfs(n *TreeNode, depth int) int {
	if n == nil {
		return 0
	}

	l := dfs(n.Left, depth+1)
	r := dfs(n.Right, depth+1)

	return max(l, r) + 1
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
