# Floyd-Warshall Algorithm Explanation

Floyd-Warshall Algorithm is used in finding shortest paths between all pairs of vertices in a weighted (directed or undirected) graph. However, it won't work for graphs with negative cycles.

# Floyd-Warshall Algorithm Implementation

1. Come up with a n by n table where n refers to the number of vertices. For each cell **[i][j]** fill it with the distance from i to j. Incase there is no direct connection, the cell value could be infinity.
2. Iterate through the table in such a way that you fill the cell **[i][j]** with the minimum distance. The minimum distance is determined by the current value of **[i][j]** or the distance of i to j through an additional vertex k (**[i][k]** + **[i][j]**).
3. Iteration continues until all possible combinations of i, j and k are considered.

# Example

```go
import "github.com/o-richard/DSA/floydwarshall"
                        .
                        .
                        .
graph := make([][]uint, 4)
graph[0] = []uint{0, 3, math.MaxInt, 5}
graph[1] = []uint{2, 0, math.MaxInt, 4}
graph[2] = []uint{math.MaxInt, 1, 0, math.MaxInt}
graph[3] = []uint{math.MaxInt, math.MaxInt,2, 0}
// Returns a graph with the shortes distances included
fmt.Println(floydwarshall.FloydWarshall(graph, 4))
```

# Applications

- To find the shortest path is a directed graph.
- To find the transitive closure of directed graphs.
- To find the Inversion of real matrices.
- For testing whether an undirected graph is bipartite (can be divided into two disjoint sets of nodes with no edges between nodes in the same set).

# Tests
