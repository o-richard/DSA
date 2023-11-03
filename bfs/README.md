# Breadth first search's Explanation

Breadth-First Search is used to tranverse data structures such as trees and graph in a systematic way.

# Breadth first search's Implementation

Visualize you are in a building with many rooms and you want to explore all rooms. A room may have 1 to n doors connecting to adjacent rooms.

1. Come up with a way to track rooms pending visitation and already visited rooms.
A [queue](../queue/README.md) may be used to track rooms pending visitation. Pick the room you are in to start with and add it to the queue.

An empty list of already visited rooms is also generated. Put the room you are in on the visited list also.

2. Explore the queue as along as it is not empty
Remove the first room in the queue. Explore unvisited adjacent rooms while adding them at the end of the queue and also adding them to the visited list.


# Example
```go
"github.com/o-richard/DSA/bfs"
                        .
                        .
                        .
graph := make(map[string][]string)
graph["0"] = []string {"1", "2", "3"}
graph["1"] = []string {"0", "2"}
graph["2"] = []string {"0", "1", "4"}
graph["3"] = []string {"0"}
graph["4"] = []string {"2"}
// Returns an array of vertices explored through Breadth First Search
fmt.Println(bfs.BFS(graph, "0"))
```

# Applications
- Path finding algorithms
- Web Crawling
- Cycle detection in an undirected graph
- Finding the minimum spanning tree

# Tests
