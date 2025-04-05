package btreezigzaglevelordertraversal

func zigzagLevelOrder(root *TreeNode) (result [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				levelNodes[i] = node.Val
			} else {
				levelNodes[levelSize-i-1] = node.Val
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelNodes)
		level++
	}

	return
}
