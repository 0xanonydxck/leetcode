package findpathexistingraph

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

func buildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := buildGraph(edges)
	visited := map[int]bool{source: true}
	q := new(Queue)
	q.Enqueue(source)

	for !q.Empty() {
		current := q.Dequeue()
		if current == destination {
			return true
		}

		for _, node := range graph[current] {
			if !visited[node] {
				visited[node] = true
				q.Enqueue(node)
			}
		}
	}

	return false
}
