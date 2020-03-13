package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}

	//On every node we will check if the target tree exists
	if isSame(s, t) {
		return true
	}
	//Target t can be found either on left of this node or right of the node, in either case if found we need to return true
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
