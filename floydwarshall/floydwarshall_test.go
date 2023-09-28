package floydwarshall_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/o-richard/DSA/floydwarshall"
)

func TestFloydWarshall(t *testing.T) {
	graph := make([][]uint, 4)
	graph[0] = []uint{0, 3, math.MaxInt, 5}
	graph[1] = []uint{2, 0, math.MaxInt, 4}
	graph[2] = []uint{math.MaxInt, 1, 0, math.MaxInt}
	graph[3] = []uint{math.MaxInt, math.MaxInt, 2, 0}
	expectedGraph := make([][]uint, 4)
	expectedGraph[0] = []uint{0, 3, 7, 5}
	expectedGraph[1] = []uint{2, 0, 6, 4}
	expectedGraph[2] = []uint{3, 1, 0, 5}
	expectedGraph[3] = []uint{5, 3, 2, 0}
	
	actualOutput, _ := floydwarshall.FloydWarshall(graph)
	if !reflect.DeepEqual(actualOutput, expectedGraph) {
		t.Errorf("expected %v but got %v for the graph %v", expectedGraph, actualOutput, graph)
	}

	badGraph := make([][]uint, 2)
	badGraph[0] = []uint{0, 3, math.MaxInt, 5}
	_, err := floydwarshall.FloydWarshall(badGraph)
	if (err == nil) != false {
		t.Errorf("expected an error since the adjacent matrix is not uniform")
	}
}
