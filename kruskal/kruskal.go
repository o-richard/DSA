package kruskal

import (
	"slices"
)

type edge struct {
	source, destination string
	weight int
}

type graph struct {
    uniqueVertices []string
	edges []edge
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

func find(childParentRelation map[string]string, child string) string {
    if childParentRelation[child] != child {
        childParentRelation[child] = find(
            childParentRelation, childParentRelation[child],
        )
    }
    return childParentRelation[child]
}

func union(childParentRelation map[string]string, nodeRankRelation map[string]int, x, y string) {
    // Attach the smaller rank under the root of the higher rank tree (the higher rank one becomes the parent of the less rank one)
    if nodeRankRelation[x] < nodeRankRelation[y] {
        childParentRelation[x] = y
    }
    if nodeRankRelation[x] > nodeRankRelation[y] {
        childParentRelation[y] = x
    }
    if nodeRankRelation[x] == nodeRankRelation[y] {
        childParentRelation[y] = x
        nodeRankRelation[x] = nodeRankRelation[x] + 1
    }
}

func (g *graph) Kruskal() ([]edge, int) {
    var result []edge
    var minCost int
    // Keep track of the indices in the original sorted edges and the result edges
    var indexOriginalEdges, indexResultEdges int

    slices.SortStableFunc(g.edges, func(a, b edge) int {
        return a.weight - b.weight
    })

    verticesCount := len(g.uniqueVertices)
    childParentRelation := make(map[string]string)
    for _, v := range g.uniqueVertices {
        childParentRelation[v] = v
    }
    nodeRankRelation := make(map[string]int)

    for indexResultEdges < verticesCount - 1 {
        minEdge := g.edges[indexOriginalEdges]
        indexOriginalEdges++

        sourceParent := find(childParentRelation, minEdge.source)
        destinationParent := find(childParentRelation, minEdge.destination)

        if sourceParent != destinationParent {
            indexResultEdges++
            result = append(result, minEdge)
            minCost += minEdge.weight
            union(childParentRelation, nodeRankRelation, sourceParent, destinationParent)
        }

    }

    return result, minCost
}
