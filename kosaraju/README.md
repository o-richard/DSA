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
fmt.Println(graph.SCC())
```