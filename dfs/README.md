# Depth first search's Explanation

Depth-First Search is used to tranverse data structures such as trees and graph in a systematic way.

# Depth first search's Implementation

Visualize you are in a building with many rooms and you want to explore all rooms. A room may have 1 to n doors connecting to adjacent rooms.

1. Choose a room to start with.
2. Explore an adjacent unvisited room, if you get to a room with no unvisited adjacent rooms, go to the most recent room with unvisited adjacent rooms.
3. Keep exploring until you have explored all rooms.

# Example

```go
"github.com/o-richard/DSA/dfs"
                        .
                        .
                        .
graph := make(map[string][]string)
graph["0"] = []string {"1", "2", "3"}
graph["1"] = []string {"0", "2"}
graph["2"] = []string {"0", "1", "4"}
graph["3"] = []string {"0"}
graph["4"] = []string {"2"}
// Returns an array of vertices explored through Depth First Search
fmt.Println(dfs.DFS(graph, "0"))
```

# Applications
- Path Finding
- Finding the strongly connected components of a graph
- Cycle detection
- Test if the graph is bipartite (can be divided into two disjoint sets of nodes with no edges between nodes in the same set).

# Tests