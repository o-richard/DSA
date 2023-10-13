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
fmt.Println(bfs.BFS(graph, "0"))
```