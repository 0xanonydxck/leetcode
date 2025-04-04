package stepbystepdirectionbtree

const (
	UP    = 'U'
	LEFT  = 'L'
	RIGHT = 'R'
)

func getDirections(root *TreeNode, startValue int, destValue int) string {
	middlePoint := findLowestAncestor(root, startValue, destValue)
	startPath := buildUpPath(middlePoint, startValue)
	destPath := buildDestPath(middlePoint, destValue)
	return startPath + destPath
}

func buildUpPath(root *TreeNode, point int) string {
	var dfs func(root *TreeNode, point int, path []rune) (bool, []rune)
	dfs = func(root *TreeNode, point int, path []rune) (bool, []rune) {
		if root == nil {
			return false, path
		}

		if root.Val == point {
			return true, path
		}

		left, leftPath := dfs(root.Left, point, append(path, UP))
		if left {
			return left, leftPath
		}

		right, rightPath := dfs(root.Right, point, append(path, UP))
		if right {
			return right, rightPath
		}

		return false, path
	}

	_, path := dfs(root, point, make([]rune, 0))
	return string(path)
}

func buildDestPath(root *TreeNode, point int) string {
	var dfs func(root *TreeNode, point int, path []rune) (bool, []rune)
	dfs = func(root *TreeNode, point int, path []rune) (bool, []rune) {
		if root == nil {
			return false, path
		}

		if root.Val == point {
			return true, path
		}

		left, leftPath := dfs(root.Left, point, append(path, LEFT))
		if left {
			return left, leftPath
		}

		right, rightPath := dfs(root.Right, point, append(path, RIGHT))
		if right {
			return right, rightPath
		}

		return false, path
	}

	_, path := dfs(root, point, make([]rune, 0))
	return string(path)
}

func findLowestAncestor(root *TreeNode, startValue, destValue int) *TreeNode {
	var dfs func(root *TreeNode, dest int, path []*TreeNode) (bool, []*TreeNode)
	dfs = func(root *TreeNode, dest int, path []*TreeNode) (bool, []*TreeNode) {
		if root == nil {
			return false, path
		}

		path = append(path, root)
		if root.Val == dest {
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

	originPath := make([]*TreeNode, 0)
	_, originPath = dfs(root, startValue, originPath)

	destPath := make([]*TreeNode, 0)
	_, destPath = dfs(root, destValue, destPath)

	var middle *TreeNode
	minLen := 0
	if len(originPath) > len(destPath) {
		minLen = len(destPath)
	} else {
		minLen = len(originPath)
	}

	for i := range minLen {
		if originPath[i].Val == destPath[i].Val {
			middle = originPath[i]
		}
	}

	return middle
}
