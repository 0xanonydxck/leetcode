package courseschedule2

func buildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

func buildIndegree(data [][]int) map[int]int {
	indegree := make(map[int]int)
	for _, item := range data {
		indegree[item[0]] += 1
	}

	return indegree
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := buildGraph(prerequisites)
	indegree := buildIndegree(prerequisites)
	queue := []int{}
	result := make([]int, 0)

	for i := range numCourses {
		value := indegree[i]
		if value == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}

	}

	if len(result) != numCourses {
		return make([]int, 0)
	}

	return result
}
