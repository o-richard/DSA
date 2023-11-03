package kruskal_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/kruskal"
)

func testEquality(a, b []kruskal.Edge) bool {
	if len(a) != len(b) {
		return true
	}

	// Store sources and destinations. Sources as keys. Destinations as values.
	mapA := make(map[string]map[string]int)
	mapB := make(map[string]map[string]int)

	for _, e := range a {
		v := mapA[e.Source]
		if v == nil {
			mapA[e.Source] = make(map[string]int)
		}
		mapA[e.Source][e.Destination] = e.Weight
	}
	for _, e := range b {
		v := mapB[e.Source]
		if v == nil {
			mapB[e.Source] = make(map[string]int)
		}
		mapB[e.Source][e.Destination] = e.Weight
	}

	return reflect.DeepEqual(mapA, mapB)
}

func TestKruskal(t *testing.T) {
	graph := kruskal.InitGraph()
	graph.AddEdge("A", "B", 4)
	graph.AddEdge("A", "D", 4)
	graph.AddEdge("B", "D", 2)
	graph.AddEdge("D", "C", 3)
	graph.AddEdge("C", "E", 3)
	graph.AddEdge("D", "E", 4)
	graph.AddEdge("D", "F", 2)
	graph.AddEdge("F", "E", 3)
	expectedEdges := []kruskal.Edge{
		{Source: "B", Destination: "D", Weight: 2},
		{Source: "D", Destination: "F", Weight: 2},
		{Source: "D", Destination: "C", Weight: 3},
		{Source: "A", Destination: "B", Weight: 4},
		{Source: "C", Destination: "E", Weight: 3},
	}
	expectedMinCost := 14
	actualEdges, actualMinCost := graph.Kruskal()

	if !testEquality(expectedEdges, actualEdges) {
		t.Errorf("expected %v as the edges but got %v", expectedEdges, actualEdges)
	}
	if expectedMinCost != actualMinCost {
		t.Errorf("expected %v as the minimum cost but got %v", expectedMinCost, actualMinCost)
	}
}
