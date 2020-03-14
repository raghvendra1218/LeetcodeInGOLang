package main

import "LeetcodeInGOLang/utils"

//Approach 1: Recursive Solution, for each node we can check if it's left sub tree is equal to the right subtree
func isSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
