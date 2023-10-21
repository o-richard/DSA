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
fmt.Println(graph.Kruskal())
```