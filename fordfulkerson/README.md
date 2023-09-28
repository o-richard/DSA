# Ford-Fulkerson Algorithm Explanation

Ford-Fulkerson Algorithm is used to solve the maximum amount of flow that can be sent from a source node to a sink node in a directed graph.

# Ford-Fulkerson Algorithm Implementation

Imagine you have a network of pipes where each pipe has a certain capacity to carry water from one place to another. You want to maximize the water flow from a source to a destination.

1. Initialization - At the start, the pipes have no water and the initial flow is 0.
2. Find a path from source to destination that has residual capacity (unfilled space in the pipes).
3. Find the maximum amount of water that can be pushed through the recently found path without exceeding the capacity of any pipe in the path.
4. Augemment the max flow by the amount found in step 3.
5. Update the network by reducing the capacity of all the pipes in the path with the max amount obtained at step 2. (This is partially filling the pipe with water).
6. Repeat 2 - 5 as along as there are no more paths from the source to destination with the available capacity.

# Example

```go
"github.com/o-richard/DSA/fordfulkerson"
                        .
                        .
                        .
graph := map[string]map[string]int {
    "S": {
        "A": 7,
        "D": 4,
    },
    "A": {
        "B": 5,
        "C": 3,
    },
    "D": {
        "A": 3,
        "C": 2,
    },
    "B": {
        "T": 8,
    },
    "C": {
        "B": 3,
        "T": 5,
    },
}
// Returns the maximum flow in the network
fmt.Println(fordfulkerson.FordFulkerson(graph, "S", "T"))
```

# Applications

- Network Flow Optimization
- Water distribution pipeline
- Bipartite matching problem
- Circulation with demands

# Tests