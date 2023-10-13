```go
"github.com/o-richard/DSA/bellmanford"
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
fmt.Println(graph.BellmanFord("A"))
```