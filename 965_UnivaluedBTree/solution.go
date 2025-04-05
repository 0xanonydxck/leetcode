package univaluedbtree

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if !visited[current.Val] {
			visited[current.Val] = true
		}

		if len(visited) > 1 {
			return false
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return true
}
