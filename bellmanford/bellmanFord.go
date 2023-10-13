package bellmanford

import (
	"errors"
	"math"

	"golang.org/x/exp/slices"
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
	distance := make(map[string]int)

	for _, v := range g.uniqueVertices {
		distance[v] = math.MaxInt
	}
	distance[source] = 0

	// Relax Edges
	for i, size := 0, len(g.uniqueVertices); i < size-1; i++ {
		for _, edge := range g.edges {
			if distance[edge.source] != math.MaxInt64 && distance[edge.source]+edge.weight < distance[edge.destination] {
				distance[edge.destination] = distance[edge.source] + edge.weight
			}
		}
	}

	// Check for negative cycles
	for _, edge := range g.edges {
		if distance[edge.source] != math.MaxInt64 && distance[edge.source]+edge.weight < distance[edge.destination] {
			return distance, errors.New("negative cycle found")
		}
	}

	// Distance from source
	return distance, nil
}
