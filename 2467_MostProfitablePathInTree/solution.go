package mostprofitablepathintree

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	// Create adjacency list representation of the tree
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	// Find Bob's path to root (node 0)
	bobPath := make(map[int]int) // node -> time
	visited := make([]bool, n)

	var findBobPath func(node, time int) bool
	findBobPath = func(node, time int) bool {
		if node == 0 {
			bobPath[node] = time
			return true
		}

		visited[node] = true
		for _, next := range graph[node] {
			if !visited[next] {
				if findBobPath(next, time+1) {
					bobPath[node] = time
					return true
				}
			}
		}
		return false
	}

	visited[bob] = true
	findBobPath(bob, 0)

	// Find Alice's maximum profit path
	maxProfit := math.MinInt32

	var dfs func(node, parent, time, profit int)
	dfs = func(node, parent, time, profit int) {
		// Calculate gate profit/cost for current node
		currAmount := amount[node]

		bobTime, exists := bobPath[node]
		if exists {
			if time == bobTime {
				currAmount /= 2 // They arrive simultaneously
			} else if time > bobTime {
				currAmount = 0 // Bob already opened the gate
			}
		}

		profit += currAmount

		// Check if it's a leaf node (only connected to parent)
		isLeaf := true
		for _, next := range graph[node] {
			if next != parent {
				isLeaf = false
				dfs(next, node, time+1, profit)
			}
		}

		if isLeaf {
			maxProfit = max(maxProfit, profit)
		}
	}

	dfs(0, -1, 0, 0)
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
