package floydwarshall

import "errors"

func FloydWarshall(graph [][]uint) ([][]uint, error) {
	verticesCount := len(graph)
	for i := range graph {
		if len(graph[i]) != verticesCount {
			return graph, errors.New("the number of size of the graph should match the number of vertices")
		}
	}
	
	for k := 0; k < verticesCount; k++ {
		for i := 0; i < verticesCount; i++ {
			for j := 0; j < verticesCount; j++ {
				graph[i][j] = min(graph[i][j], graph[i][k] + graph[k][j])
			}
		}
	}

	return graph, nil
}