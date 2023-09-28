package dfs_test

import (
	"reflect"
	"testing"

	"github.com/o-richard/DSA/dfs"
)

func TestDFS(t *testing.T) {
	graph := make(map[string][]string)
	graph["0"] = []string{"1", "2", "3"}
	graph["1"] = []string{"0", "2"}
	graph["2"] = []string{"0", "1", "4"}
	graph["3"] = []string{"0"}
	graph["4"] = []string{"2"}

	tests := map[string]struct {
		source string
		output []string
	}{
		"Test 1": {source: "0", output: []string{"0", "1", "2", "4", "3"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualOutput := dfs.DFS(graph, test.source)
			if !reflect.DeepEqual(test.output, actualOutput) {
				t.Errorf("expected %v but obtained %v for the source '%v' for the graph %v", test.output, actualOutput, test.source, graph)
			}
		})
	}
}
