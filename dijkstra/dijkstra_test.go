package dijkstra_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/dijkstra"
)

func TestDijkstra(t *testing.T) {
	graph := map[string]map[string]int{
		"A": {
			"B": 4,
			"C": 5,
		},
		"C": {
			"A": 5,
			"B": 11,
			"E": 3,
		},
		"B": {
			"A": 4,
			"C": 11,
			"D": 9,
			"E": 7,
		},
		"D": {
			"B": 9,
			"E": 13,
			"F": 2,
		},
		"E": {
			"C": 3,
			"B": 7,
			"D": 13,
			"F": 6,
		},
		"F": {
			"D": 2,
			"E": 6,
		},
	}

	source := "A"
	expectedOutput := map[string]int{
		"A": 0,
		"B": 4,
		"C": 5,
		"D": 13,
		"E": 8,
		"F": 14,
	}
	actualOutput := dijkstra.Dijkstra(graph, source)

	if !reflect.DeepEqual(expectedOutput, actualOutput) {
		t.Errorf("expected the answer %v but got %v", expectedOutput, actualOutput)
	}
}
