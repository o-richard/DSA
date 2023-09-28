package kosaraju_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/kosaraju"
)

func testEquality(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	mapB := make(map[int][][]string)
	for i := range b {
		length := len(b[i])
		mapB[length] = append(mapB[length], b[i])
	}

	for _, current := range a {
		length := len(current)
		slices, ok := mapB[length]
		if !ok {
			return false
		}

		mapCurrent := make(map[string]int, length)
		for i := range current {
			mapCurrent[current[i]]++
		}

		var found bool
	compare:
		for i, slice := range slices {
			mapCompare := make(map[string]int, length)
			for j := range slice {
				mapCompare[slice[j]]++
			}
			if !reflect.DeepEqual(mapCurrent, mapCompare) {
				continue
			}
			found = true
			mapB[length] = append(mapB[length][:i], mapB[length][i+1:]...)
			if len(mapB[length]) == 0 {
				delete(mapB, length)
			}
			break compare
		}
		if !found {
			return false
		}
	}

	return true
}

func TestSCC(t *testing.T) {
	graph := kosaraju.InitGraph()
	graph.AddEdge("0", "1")
	graph.AddEdge("1", "2")
	graph.AddEdge("2", "3")
	graph.AddEdge("2", "4")
	graph.AddEdge("3", "0")
	graph.AddEdge("4", "5")
	graph.AddEdge("5", "6")
	graph.AddEdge("6", "4")
	graph.AddEdge("6", "7")

	expectedOutput := [][]string{
		{"0", "3", "2", "1"},
		{"4", "5", "6"},
		{"7"},
	}
	actualOutput := graph.SCC()
	if !testEquality(expectedOutput, actualOutput) {
		t.Errorf("expected %v but got %v", expectedOutput, actualOutput)
	}
}
