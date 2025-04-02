package reachablenodesrestriction

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MinInt  = -1 << (intSize - 1)    // MinInt32 or MinInt64 depending on intSize.
)

type Queue []int

func (q *Queue) Enqueue(n int) {
	*q = append(*q, n)
}

func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		return MinInt
	}

	value := (*q)[0]
	*q = (*q)[1:]
	return value
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func buildGraph(n int, edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	restrictedSet := make(map[int]bool)
	for _, r := range restricted {
		restrictedSet[r] = true
	}

	graph := buildGraph(n, edges)
	if restrictedSet[0] {
		return 0
	}

	visited := make(map[int]bool)
	q := new(Queue)
	q.Enqueue(0)
	visited[0] = true
	result := 0

	for !q.Empty() {
		current := q.Dequeue()
		result++

		if neighbors, exists := graph[current]; exists {
			for _, node := range neighbors {
				if !visited[node] && !restrictedSet[node] {
					visited[node] = true
					q.Enqueue(node)
				}
			}
		}
	}

	return result
}
