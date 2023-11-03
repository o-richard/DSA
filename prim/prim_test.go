package prim_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/prim"
)

func testEquality(a, b []prim.Edge) bool {
	if len(a) != len(b) {
		return true
	}

	// Store sources and destinations. Sources as keys. Destinations as values.
	mapA := make(map[string]map[string]int)
	mapB := make(map[string]map[string]int)

	// NOTE: The test uses a weighted undirected graph.
	for _, e := range a {
		if v := mapA[e.Source]; v == nil {
			mapA[e.Source] = make(map[string]int)
		}
		if v := mapA[e.Destination]; v == nil {
			mapA[e.Destination] = make(map[string]int)
		}

		mapA[e.Source][e.Destination] = e.Weight
		mapA[e.Destination][e.Source] = e.Weight
	}
	for _, e := range b {
		if v := mapB[e.Source]; v == nil {
			mapB[e.Source] = make(map[string]int)
		}
		if v := mapB[e.Destination]; v == nil {
			mapB[e.Destination] = make(map[string]int)
		}
		mapB[e.Source][e.Destination] = e.Weight
		mapB[e.Destination][e.Source] = e.Weight
	}
	
	return reflect.DeepEqual(mapA, mapB)
}


func TestPrim(t *testing.T) {
	graph := map[string]map[string]int {
		"A": {
			"B": 7,
			"C": 4,
		},
		"B": {
			"A": 7,
			"C": 2,
		},
		"C": {
			"A": 4,
			"B": 2,
			"F": 4,
			"D": 3,
			"E": 2,
		},
		"D": {
			"F": 9,
			"C": 3,
		},
		"E": {
			"F": 3,
			"C": 2,
		},
		"F": {
			"E": 3,
			"D": 9,
			"C": 4,
		},
	}
	expectedEdges := []prim.Edge{
		{Source: "A", Destination: "C", Weight: 4},
		{Source: "B", Destination: "C", Weight: 2},
		{Source: "C", Destination: "D", Weight: 3},
		{Source: "C", Destination: "E", Weight: 2},
		{Source: "E", Destination: "F", Weight: 3},
	}
	expectedMinCost := 14
	actualEdges, actualMinCost := prim.Prim(graph)

	if !testEquality(expectedEdges, actualEdges) {
		t.Errorf("expected %v as the edges but got %v", expectedEdges, actualEdges)
	}
	if expectedMinCost != actualMinCost {
		t.Errorf("expected %v as the minimum cost but got %v", expectedMinCost, actualMinCost)
	}
}
