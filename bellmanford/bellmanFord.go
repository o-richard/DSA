package bellmanford

import (
	"errors"
	"math"
	"slices"
)

type edge struct {
	source, destination string
	weight              int
}

type graph struct {
	uniqueVertices []string
	edges          []edge
}

func InitGraph() graph {
	return graph{}
}

func (g *graph) AddEdge(source, destination string, weight int) {
	if !slices.Contains(g.uniqueVertices, source) {
		g.uniqueVertices = append(g.uniqueVertices, source)
	}
	if !slices.Contains(g.uniqueVertices, destination) {
		g.uniqueVertices = append(g.uniqueVertices, destination)
	}
	newEdge := edge{source: source, destination: destination, weight: weight}
	g.edges = append(g.edges, newEdge)
}

func (g *graph) BellmanFord(source string) (map[string]int, error) {
	// Store the distance from the provided source to a vertex
	distance := make(map[string]int)
	for _, v := range g.uniqueVertices {
		distance[v] = math.MaxInt
	}
	distance[source] = 0

	// Relax Edges |V| - 1 times where V refers to number of unique vertices
	for i, size := 0, len(g.uniqueVertices); i < size-1; i++ {
		for _, edge := range g.edges {
			if distance[edge.source] != math.MaxInt && distance[edge.source]+edge.weight < distance[edge.destination] {
				distance[edge.destination] = distance[edge.source] + edge.weight
			}
		}
	}

	// Check for negative cycles. (If a value changes, there is a negative cycle)
	for _, edge := range g.edges {
		if distance[edge.source] != math.MaxInt && distance[edge.source]+edge.weight < distance[edge.destination] {
			return distance, errors.New("negative cycle found")
		}
	}

	return distance, nil
}
