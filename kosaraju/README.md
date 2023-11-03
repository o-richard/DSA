# Kosaraju's Algorithm Explanation

Kosaraju's Algorithm is used to find Strongly Connected Components in a directed graph.

# Kosaraju's Algorithm Implementation

Imagine you are in a small neighbourhood with one way paths and you would like to find which houses are directly connected to each other (You can get to any house from any house within the group, groups can be isolated).

1. Mark the houses in the order as you visit them. You start with a house, you follow the street till an endpoint and you mark the endpoint, backtrack and keep marking unmarked houses along the way.
2. After all houses have been marked, reverse the direction of all streets in the neighourhood in the reverse direction.
3. Starting from a marked house, you follow the reversed streets while marking all the houses you can reach in the path. In case you come across an already marked one, you have a Strongly Connected Component. You repeat the process for all the unmarked houses.

# Example

```go
import "github.com/o-richard/DSA/kosaraju"
                        .
                        .
                        .
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
// Obtain a slice of strongly connected components.
// Each strongly connected component has an array of vertices present.
fmt.Println(graph.SCC())
```

# Applications

- Identification of Strongly Connected Components.
- Transportation and route planning

# Tests

