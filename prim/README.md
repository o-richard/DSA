# Prim's Algorithm Explanation

Prim's Algorithm is used to find the minimum spanning tree in a graph (A subset of the edges of graph that connects all the vertices with the minimum possible edge weight without forming any cycle).

# Prim's Algorithm Implementation

Imagine you have a map of cities and you want to build a road network to connect cities while minimizing length.

1. Randomly select a single city as your starting point
2. Identify a city connected to your chosen city with the shortest road distance and it is not yet chosen
3. Continue the selction for the different cities and roads

# Example

```go
"github.com/o-richard/DSA/prim"
                .
                .
                .
graph := map[string]map[string]int {
    "0": {
        "1": 4,
        "7": 8,
    },
    "1": {
        "0": 4,
        "7": 11,
        "2": 8,
    },
    "7": {
        "0": 8,
        "1": 11,
        "8": 7,
        "6": 1,
    },
    "2": {
        "1": 8,
        "3": 7,
        "8": 2,
        "5": 4,
    },
    "8": {
        "2": 2,
        "6": 6,
        "7": 7,
    },
    "6": {
        "7": 1,
        "8": 6,
        "5": 2,
    },
    "3": {
        "2": 7,
        "4": 9,
        "5": 14,
    },
    "5": {
        "6": 2,
        "2": 4,
        "3": 14,
        "4": 10,
    },
    "4": {
        "5": 10,
        "3": 9,
    },
}
fmt.Println(prim.Prim(graph))
```

# Applications

- Find the minimum spanning tree
- Network design
- Electrical wiring

# Tests

