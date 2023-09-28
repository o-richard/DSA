package fordfulkerson

import (
	"math"
)

func bfs(graph map[string]map[string]int, source, sink string) (map[string]string, bool) {
	// Store the child as a key and its parent as a value.
	childParentRelation := make(map[string]string)

	visited := map[string]bool{source: true}
	queue := []string{source}

	// As along as the queue is not empty
	for queueLength := len(queue); queueLength > 0; {
		// Dequeue
		vertex := queue[0]
		queue = queue[1:]
		queueLength--

		for adjacent := range graph[vertex] {
			// The path should be unvistited and has a residual capacity (greater than zero)
			if !visited[adjacent] && graph[vertex][adjacent] > 0 {
				visited[adjacent] = true
				queue = append(queue, adjacent)
				queueLength++

				childParentRelation[adjacent] = vertex

				if adjacent == sink {
					return childParentRelation, true
				}

			}
		}
	}

	return childParentRelation, false
}

func FordFulkerson(graph map[string]map[string]int, source, sink string) int {
	var maxFlow int

	for childParentRelation, isPath := bfs(graph, source, sink); isPath; childParentRelation, isPath = bfs(graph, source, sink) {
		pathFlow := math.MaxInt

		// Find the minimum residual capacity of the edges in the augmenting path i.e max flow through the path
		for s := sink; s != source; s = childParentRelation[s] {
			pathFlow = min(pathFlow, graph[childParentRelation[s]][s])
		}

		maxFlow += pathFlow

		// Update residual capacities of the edges
		for v := sink; v != source; v = childParentRelation[v] {
			// Decrease capacity of the forward edge
			graph[childParentRelation[v]][v] -= pathFlow
		}

	}

	return maxFlow
}
