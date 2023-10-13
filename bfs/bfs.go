package bfs

import "golang.org/x/exp/slices"

func BFS(graph map[string][]string, start string) []string {
	queue := []string{start}
	visited := []string{start}

	for queueLength := len(queue); queueLength > 0; {
		vertex := queue[0]
		queue = queue[1:]
		queueLength--

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