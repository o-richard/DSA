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
fmt.Println(fordfulkerson.FordFulkerson(graph, "S", "T"))
```