package kosaraju

import (
	"golang.org/x/exp/slices"
)

type graph struct {
	uniqueVertices []string
	graph          map[string][]string
}

func InitGraph() graph {
	return graph{graph: make(map[string][]string)}
}

func (g *graph) AddEdge(source, destination string) {
	if !slices.Contains(g.uniqueVertices, source) {
		g.uniqueVertices = append(g.uniqueVertices, source)
	}
	if !slices.Contains(g.uniqueVertices, destination) {
		g.uniqueVertices = append(g.uniqueVertices, destination)
	}
	g.graph[source] = append(g.graph[source], destination)
}

func (g *graph) fillOrder(vertex string, visitedVertices map[string]bool, stack *[]string) {
	visitedVertices[vertex] = true
	for _, destVertex := range g.graph[vertex] {
		if !visitedVertices[destVertex] {
			g.fillOrder(destVertex, visitedVertices, stack)
		}
	}
	*stack = append(*stack, vertex)
}

func (g *graph) transpose() graph {
	reversedGraph := InitGraph()

	for source, dests := range g.graph {
		for _, dest := range dests {
			reversedGraph.AddEdge(dest, source)
		}
	}

	return reversedGraph
}

func (g *graph) dfs(vertex string, visitedVertices map[string]bool, sample []string) []string {
	visitedVertices[vertex] = true
	sample = append(sample, vertex)
	for _, destVertex := range g.graph[vertex] {
		if !visitedVertices[destVertex] {
			sample = g.dfs(destVertex, visitedVertices, sample)
		}
	}
	return sample
}

func (g *graph) SCC() [][]string {
	var stack []string
	visitedVertices := make(map[string]bool)

	for _, v := range g.uniqueVertices {
		if !visitedVertices[v] {
			g.fillOrder(v, visitedVertices, &stack)
		}
	}

	reversedGraph := g.transpose()

	for _, v := range g.uniqueVertices {
		visitedVertices[v] = false
	}

	var scc [][]string

	for stackLength := len(stack); stackLength != 0; stackLength-- {
		vertex := stack[stackLength-1]
		stack = stack[:stackLength-1]

		if !visitedVertices[vertex] {
			var sample []string
			sample = reversedGraph.dfs(vertex, visitedVertices, sample)
			scc = append(scc, sample)
		}
	}

	return scc
}
