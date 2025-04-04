package lowestcommonancestorbtree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(root, dest *TreeNode, path []*TreeNode) (bool, []*TreeNode)
	dfs = func(root, dest *TreeNode, path []*TreeNode) (bool, []*TreeNode) {
		if root == nil {
			return false, path
		}

		path = append(path, root)
		if root.Val == dest.Val {
			return true, path
		}

		left, leftPath := dfs(root.Left, dest, path)
		if left {
			return left, leftPath
		}

		right, rightPath := dfs(root.Right, dest, path)
		if right {
			return right, rightPath
		}

		return false, path
	}

	pPath := make([]*TreeNode, 0)
	_, pPath = dfs(root, p, pPath)

	qPath := make([]*TreeNode, 0)
	_, qPath = dfs(root, q, qPath)

	var lowest *TreeNode
	minLen := 0
	if len(pPath) > len(qPath) {
		minLen = len(qPath)
	} else {
		minLen = len(pPath)
	}

	for i := range minLen {
		if pPath[i].Val == qPath[i].Val {
			lowest = pPath[i]
		}
	}

	return lowest
}
