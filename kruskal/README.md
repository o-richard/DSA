# Kruskal's Algorithm Explanation

Kruskal's Algorithm is used to find the minimum spanning tree in a graph (A subset of the edges of graph that connects all the vertices with the minimum possible edge weight without forming any cycle).

# Kruskal's Algorithm Implementation

Imagine you have a map of cities and you want to build a road network to connect cities while minimizing length.

1. Sort all possible roads on the map in an ascending order based on the length.
2. Pick the shortest road connecting two cities.
3. While picking out the shortest road, you ensure there are no cycles in the road network (there are no closed paths). If there would be a cycle, the next shortest edge is picked.
4. Repeat for all roads and one obtains the minimum length connecting cities without any loops.

# Example

```go
"github.com/o-richard/DSA/kruskal"
                .
                .
                .
graph := kruskal.InitGraph()
graph.AddEdge("0", "1", 10)
graph.AddEdge("0", "2", 6)
graph.AddEdge("0", "3", 5)
graph.AddEdge("1", "3", 15)
graph.AddEdge("2", "3", 4)
// Returns the edges in the minimum spanning tree and the minimum cost
fmt.Println(graph.Kruskal())
```

# Applications

- Find the minimum spanning tree
- Network design
- Electrical wiring

# Tests
