package bellmanford_test

import (
	"testing"

	"github.com/o-richard/DSA/bellmanford"
)

type edge struct {
	source, destination string
	weight              int
}
type args struct {
	sourceInput     string
	distanceOutput  map[string]int
	isSuccessOutput bool
}

func TestBellmanFord(t *testing.T) {
	tests := map[string]struct {
		edges []edge
		args  []args
	}{
		"Negative Cycles Absent": {
			edges: []edge{
				{source: "A", destination: "B", weight: 4},
				{source: "A", destination: "C", weight: 2},
				{source: "B", destination: "C", weight: 3},
				{source: "C", destination: "B", weight: 1},
				{source: "B", destination: "D", weight: 2},
				{source: "C", destination: "E", weight: 5},
				{source: "B", destination: "E", weight: 3},
				{source: "C", destination: "D", weight: 4},
				{source: "E", destination: "D", weight: -5},
			},
			args: []args{
				{
					sourceInput: "A", isSuccessOutput: true,
					distanceOutput: map[string]int{"A": 0, "B": 3, "C": 2, "D": 1, "E": 6},
				},
				{
					sourceInput: "B", isSuccessOutput: true,
					distanceOutput: map[string]int{"B": 0, "C": 3, "D": -2, "E": 3},
				},
				{
					sourceInput: "C", isSuccessOutput: true,
					distanceOutput: map[string]int{"B": 1, "C": 0, "D": -1, "E": 4},
				},
				{sourceInput: "F", distanceOutput: map[string]int{"F": 0}, isSuccessOutput: true},
			},
		},
		"Negative Cycles Present": {
			edges: []edge{
				{source: "A", destination: "B", weight: 4},
				{source: "A", destination: "C", weight: 2},
				{source: "B", destination: "C", weight: 3},
				{source: "C", destination: "B", weight: 1},
				{source: "B", destination: "D", weight: 2},
				{source: "D", destination: "B", weight: 7},
				{source: "C", destination: "E", weight: 5},
				{source: "E", destination: "C", weight: -2},
				{source: "B", destination: "E", weight: 3},
				{source: "C", destination: "D", weight: 4},
				{source: "E", destination: "D", weight: -5},
				{source: "D", destination: "E", weight: -5},
			},
			args: []args{
				{sourceInput: "A"},
				{sourceInput: "B"},
				{sourceInput: "F", distanceOutput: map[string]int{"F": 0}, isSuccessOutput: true},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			graph := bellmanford.InitGraph()
			for _, v := range test.edges {
				graph.AddEdge(v.source, v.destination, v.weight)
			}
			for _, v := range test.args {
				actualDistanceOutput, err := graph.BellmanFord(v.sourceInput)
				for k, distance := range v.distanceOutput {
					if actualDistanceOutput[k] != distance {
						t.Errorf("expected the distance from the source '%v' to the destination '%v' to be %v", v.sourceInput, k, distance)
					}
				}
				if (err == nil) != v.isSuccessOutput {
					if v.isSuccessOutput {
						t.Errorf("expected no error when obtaining the distances output. Negative cycles absent in the graph.")
					} else {
						t.Errorf("expected an error when obtaining the distances output. Negative cycles present in the graph.")
					}
				}
			}
		})
	}
}
