package dijkstra

import (
	"math"

	"golang.org/x/exp/slices"
)

func Dijkstra(graph map[string]map[string]int, source string) map[string]int {
	var uniqueVertices []string
	for key, values := range graph {
		if !slices.Contains(uniqueVertices, key) {
			uniqueVertices = append(uniqueVertices, key)
		}
		for value := range values {
			if !slices.Contains(uniqueVertices, value) {
				uniqueVertices = append(uniqueVertices, value)
			}	
		}
	}

	visited := make(map[string]bool)
	distance := make(map[string]int)

	for _, v := range uniqueVertices {
		distance[v] = math.MaxInt
	}
	distance[source] = 0
	visited[source] = true

	for range uniqueVertices[1:] {
		minDistance := math.MaxInt
		var nextVertex string
		for _, vertex := range uniqueVertices {
			if !visited[vertex] && distance[vertex] < minDistance {
				minDistance = distance[vertex]
				nextVertex = vertex
			}
		}

		visited[nextVertex] = true

		for _, vertex := range uniqueVertices {
			if !visited[vertex] {
				distance[vertex] = min(distance[vertex], minDistance + graph[nextVertex][vertex])
			}
		}
	}

	return distance
}
