package bfs

import "slices"

func BFS(graph map[string][]string, start string) []string {
	queue := []string{start}
	visited := []string{start}

	// As along as the queue is not empty.
	for queueLength := len(queue); queueLength > 0; {
		// Deque
		vertex := queue[0]
		queue = queue[1:]
		queueLength--

		// Enqueue unvisited neighbours of the 'vertex'
		for _, adjacent := range graph[vertex] {
			if !slices.Contains(visited, adjacent) {
				visited = append(visited, adjacent)
				queue = append(queue, adjacent)
				queueLength++
			}
		}
	}

	return visited
}