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
fmt.Println(dijkstra.Dijkstra(graph, "A"))
```