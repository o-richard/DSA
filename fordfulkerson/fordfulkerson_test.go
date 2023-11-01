package fordfulkerson_test

import (
	"testing"

	"github.com/o-richard/DSA/fordfulkerson"
)


func TestFordFulkerson(t *testing.T) {
	graph := map[string]map[string]int {
		"S": {
			"A": 7,
			"D": 4,
		},
		"A": {
			"B": 5,
			"C": 3,
		},
		"D": {
			"A": 3,
			"C": 2,
		},
		"B": {
			"T": 8,
		},
		"C": {
			"B": 3,
			"T": 5,
		},
	}
	expectedOutput := 10
	actualOutput := fordfulkerson.FordFulkerson(graph, "S", "T")

	if expectedOutput != actualOutput {
		t.Errorf("expected %v but got %v for the graph %v", expectedOutput, actualOutput, graph)
	}
}