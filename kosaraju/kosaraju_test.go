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

	// Store the length of slices of slices as keys
	mapA := make(map[int][][]string)
	mapB := make(map[int][][]string)

	for _, v := range a {
		sliceLength := len(v)
		mapA[sliceLength] = append(mapA[sliceLength], v)
	}
	for _, v := range b {
		sliceLength := len(v)
		mapB[sliceLength] = append(mapB[sliceLength], v)
	}

	if len(mapA) != len(mapB) {
		return false
	}

	for k, v := range mapB {
		current, ok := mapA[k]
		if !ok {
			return false
		}

		if len(v) != len(current) {
			return false
		}

		for _, curr := range current {
			mapCurr := make(map[string]int)
			for i := range curr {
				mapCurr[curr[i]]++
			}
		
			var isMatch bool

			for i := range v {
				mapElem := make(map[string]int)
				for _, elem := range v[i] {
					mapElem[elem]++
				}

				if len(mapCurr) != len(mapElem) {
					continue
				}

				if !reflect.DeepEqual(mapCurr, mapElem) {
					continue
				}

				isMatch = true
				v = append(v[:i], v[i + 1:]...)
			}

			if !isMatch {
				return false
			}
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
