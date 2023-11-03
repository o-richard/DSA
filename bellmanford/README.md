# Bellman Ford's Explanation

Bellman Ford's algorithm is useful for finding the shortest paths from a single source vertex to all other vertices in a weighted, directed graph. It supports negative weights unlike Dijkstra's.


# Bellman Ford's Implementation

Visualize planning a cost effective trip to various cities. Each city will represent a vertex on a graph while costs from a city to another represents a weight on a graph.

1. Setup intial budget
For each city, initialize the cost to a high value to every other city except for the starting city (source), you set it to zero.

2. Updating budgets.

Go through all cities iteratively while updating the budget from your current position.
From your current position, you look at the total cost to reach that city from the source. If it is less than the current budget for that place, update the budget accordingly.

3. Detecting negative cycles.

Once you have updated the budgest for all cities, iterate through all cities. Incase a city's budget can be further reduced, there is a negative cycle (In theory, it will reduce your expenses infinitely as you travel).

# Example
```go

import "github.com/o-richard/DSA/bellmanford"
                        .
                        .
                        .
                        
graph := bellmanford.InitGraph()
graph.AddEdge("A", "B", 4)
graph.AddEdge("A", "C", 2)
graph.AddEdge("B", "C", 3)
graph.AddEdge("C", "B", 1)
graph.AddEdge("B", "D", 2)
graph.AddEdge("B", "E", 4)
graph.AddEdge("C", "E", 5)
graph.AddEdge("C", "D", 3)
graph.AddEdge("E", "D", -5)
// Returns the destinations and the respective distances
fmt.Println(graph.BellmanFord("A"))

```

# Applications
- Finding the shortest path in routing algorithsms.

# Tests

Tests cover two graphs. One graph has negative cycles while the other one lacks.
