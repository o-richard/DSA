package dfs

import "slices"

func dfsAlgo(graph map[string][]string, start string, visited []string) []string{
	visited = append(visited, start)
	
	unvisited := graph[start]
	for _, v := range visited {
		valIndex := slices.Index(unvisited, v)
		if valIndex != -1 {
			unvisited = append(unvisited[:valIndex], unvisited[valIndex + 1:]...)
		}
	}

	for _, next := range unvisited {
		if !slices.Contains(visited, next) {
			visited = dfsAlgo(graph, next, visited)
		}
	}

	return visited
}

func DFS(graph map[string][]string, start string) []string{
	var visited []string
	visited = dfsAlgo(graph, start, visited)
	return visited
}