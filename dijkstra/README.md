# Dijkstra's Algorithm Explanation

Dijkstra's algorithm is useful for finding the shortest paths from a single source vertex to all other vertices in a weighted, directed graph. It does not support negative weights unlike Bellman Ford's.

# Dijkstra's Algorithm Implementation

Imagine you have a map of various cities and you would like to travel from your current location to all other cities with minimal travel distance.

1. Initialization
Set all distances to other cities as a high value (infinity) while setting the distance to the source as zero.

You can use a [min heap](../heap/README.md) to keep track of cities and their shortest known distance. Add the source city into the min heap with the distance zero.

2. Iterate as along as the min heap is not empty.

Obtain the unvisited city with the shortest known distance from the heap.

Explore the neighbours of the city. Update the distance incase the total distance to reach the neighbouring city from current city is lesser than distance previously recorded for the neighbour. If the distance is updated, add the neighbouring city and new distance to the min heap.

# Example

```go
"github.com/o-richard/DSA/dijkstra"
                        .
                        .
                        .
graph := map[string]map[string]int {
    "A": {
        "B": 4,
        "C": 5,
    },
    "C": {
        "A": 5,
        "B": 11,
        "E": 3,
    },
    "B": {
        "A": 4,
        "C": 11,
        "D": 9,
        "E": 7,
    },
    "D": {
        "B": 9,
        "E": 13,
        "F": 2,
    },
    "E": {
        "C": 3,
        "B": 7,
        "D": 13,
        "F": 6,
    },
    "F": {
        "D": 2,
        "E": 6,
    },
}
// Returns the destinations and the respective distances
fmt.Println(dijkstra.Dijkstra(graph, "A"))
```

# Applications
- Finding the shortest path in routing algorithsms.

# Tests