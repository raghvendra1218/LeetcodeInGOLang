package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Approach 1: Recursive Solution, for each node we can check if it's left sub tree is equal to the right subtree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
